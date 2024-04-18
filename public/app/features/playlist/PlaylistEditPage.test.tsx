import 'whatwg-fetch';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { History, Location } from 'history';
import React from 'react';
import { type match } from 'react-router-dom';
import { TestProvider } from 'test/helpers/TestProvider';

import { locationService } from '@grafana/runtime';
import { RouteDescriptor } from 'app/core/navigation/types';
import { backendSrv } from 'app/core/services/backend_srv';

import { PlaylistEditPage, RouteParams } from './PlaylistEditPage';
import { Playlist } from './types';

jest.mock('@grafana/runtime', () => ({
  ...jest.requireActual('@grafana/runtime'),
  getBackendSrv: () => backendSrv,
}));

jest.mock('app/core/components/TagFilter/TagFilter', () => ({
  TagFilter: () => {
    return <>mocked-tag-filter</>;
  },
}));

async function getTestContext({ name, interval, items, uid }: Partial<Playlist> = {}) {
  jest.clearAllMocks();
  const playlist = { name, items, interval, uid } as unknown as Playlist;
  const queryParams = {};
  const route = {} as RouteDescriptor;
  const match = { params: { uid: 'foo' } } as unknown as match<RouteParams>;
  const location = {} as Location;
  const history = {} as History;
  const getMock = jest.spyOn(backendSrv, 'get');
  const putMock = jest.spyOn(backendSrv, 'put').mockImplementation(() => Promise.resolve());

  getMock.mockResolvedValue({
    name: 'Test Playlist',
    interval: '5s',
    items: [{ title: 'First item', type: 'dashboard_by_uid', order: 1, value: '1' }],
    uid: 'foo',
  });

  const { rerender } = render(
    <TestProvider>
      <PlaylistEditPage queryParams={queryParams} route={route} match={match} location={location} history={history} />
    </TestProvider>
  );
  await waitFor(() => expect(getMock).toHaveBeenCalledTimes(1));

  return { playlist, rerender, putMock };
}

describe('PlaylistEditPage', () => {
  describe('when mounted', () => {
    it('then it should load playlist and header should be correct', async () => {
      await getTestContext();

      expect(await screen.findByRole('heading', { name: /edit playlist/i })).toBeInTheDocument();
      expect(screen.getByRole('textbox', { name: /playlist name/i })).toHaveValue('Test Playlist');
      expect(screen.getByRole('textbox', { name: /playlist interval/i })).toHaveValue('5s');
      expect(screen.getAllByRole('row')).toHaveLength(1);
    });
  });

  describe('when submitted', () => {
    it('then correct api should be called', async () => {
      const { putMock } = await getTestContext();

      expect(await screen.findByRole('heading', { name: /edit playlist/i })).toBeInTheDocument();
      expect(locationService.getLocation().pathname).toEqual('/');
      await userEvent.clear(screen.getByRole('textbox', { name: /playlist name/i }));
      await userEvent.type(screen.getByRole('textbox', { name: /playlist name/i }), 'A Name');
      await userEvent.clear(screen.getByRole('textbox', { name: /playlist interval/i }));
      await userEvent.type(screen.getByRole('textbox', { name: /playlist interval/i }), '10s');
      fireEvent.submit(screen.getByRole('button', { name: /save/i }));
      await waitFor(() => expect(putMock).toHaveBeenCalledTimes(1));
      expect(putMock).toHaveBeenCalledWith('/api/playlists/foo', {
        uid: 'foo',
        name: 'A Name',
        interval: '10s',
        items: [{ title: 'First item', type: 'dashboard_by_uid', order: 1, value: '1' }],
      });
      expect(locationService.getLocation().pathname).toEqual('/playlists');
    });
  });
});
