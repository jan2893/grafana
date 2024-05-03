import { css } from '@emotion/css';
import React, { useState } from 'react';
import { useDebounce, useDeepCompareEffect } from 'react-use';

import { dateTime, GrafanaTheme2 } from '@grafana/data';
import { Alert, Badge, LoadingPlaceholder, useStyles2 } from '@grafana/ui';
import { MatcherFieldValue } from 'app/features/alerting/unified/types/silence-form';
import { matcherFieldToMatcher } from 'app/features/alerting/unified/utils/alertmanager';
import { MATCHER_ALERT_RULE_UID } from 'app/features/alerting/unified/utils/constants';
import { AlertmanagerAlert, Matcher, MatcherOperator } from 'app/plugins/datasource/alertmanager/types';

import { alertmanagerApi } from '../../api/alertmanagerApi';
import { isNullDate } from '../../utils/time';
import { AlertLabels } from '../AlertLabels';
import { DynamicTable, DynamicTableColumnProps, DynamicTableItemProps } from '../DynamicTable';

import { AmAlertStateTag } from './AmAlertStateTag';

interface Props {
  amSourceName: string;
  matchers: MatcherFieldValue[];
  ruleUid?: string;
}

const useDebouncedDeepCompare = (cb: () => void, debounceMs: number, dependencies: unknown[]) => {
  const [state, setState] = useState<unknown[]>();

  useDebounce(cb, debounceMs, [state]);

  useDeepCompareEffect(() => {
    setState(dependencies);
  }, [dependencies]);
};

export const SilencedInstancesPreview = ({ amSourceName, matchers: inputMatchers, ruleUid }: Props) => {
  const matchers: Matcher[] = [
    ...(ruleUid ? [{ name: MATCHER_ALERT_RULE_UID, value: ruleUid, operator: MatcherOperator.equal }] : []),
    ...inputMatchers,
  ].map(matcherFieldToMatcher);
  const useLazyQuery = alertmanagerApi.endpoints.getAlertmanagerAlerts.useLazyQuery;
  const styles = useStyles2(getStyles);
  const columns = useColumns();

  // By default the form contains an empty matcher - with empty name and value and = operator
  // We don't want to fetch previews for empty matchers as it results in all alerts returned
  const hasValidMatchers = ruleUid || inputMatchers.some((matcher) => matcher.value && matcher.name);

  const [getAlertmanagerAlerts, { currentData: alerts = [], isFetching, isError }] = useLazyQuery();

  useDebouncedDeepCompare(
    () => {
      if (hasValidMatchers || ruleUid) {
        getAlertmanagerAlerts({ amSourceName, filter: { matchers } });
      }
    },
    500,
    [amSourceName, matchers]
  );

  const tableItemAlerts = alerts.map<DynamicTableItemProps<AlertmanagerAlert>>((alert) => ({
    id: alert.fingerprint,
    data: alert,
  }));

  return (
    <div>
      <h4 className={styles.title}>
        Affected alert rule instances
        {tableItemAlerts.length > 0 ? (
          <Badge className={styles.badge} color="blue" text={tableItemAlerts.length} />
        ) : null}
      </h4>
      {!hasValidMatchers && <span>Add a valid matcher to see affected alerts</span>}
      {isError && (
        <Alert title="Preview not available" severity="error">
          Error occured when generating preview of affected alerts. Are your matchers valid?
        </Alert>
      )}
      {isFetching && <LoadingPlaceholder text="Loading affected alert rule instances..." />}
      {!isFetching && !isError && hasValidMatchers && (
        <div className={styles.table}>
          {tableItemAlerts.length > 0 ? (
            <DynamicTable
              items={tableItemAlerts}
              isExpandable={false}
              cols={columns}
              pagination={{ itemsPerPage: 10 }}
            />
          ) : (
            <span>No firing alert instances found</span>
          )}
        </div>
      )}
    </div>
  );
};

function useColumns(): Array<DynamicTableColumnProps<AlertmanagerAlert>> {
  const styles = useStyles2(getStyles);

  return [
    {
      id: 'state',
      label: 'State',
      renderCell: function renderStateTag({ data }) {
        return <AmAlertStateTag state={data.status.state} />;
      },
      size: '120px',
      className: styles.stateColumn,
    },
    {
      id: 'labels',
      label: 'Labels',
      renderCell: function renderName({ data }) {
        return <AlertLabels labels={data.labels} size="sm" />;
      },
      size: 'auto',
    },
    {
      id: 'created',
      label: 'Created',
      renderCell: function renderSummary({ data }) {
        return <>{isNullDate(data.startsAt) ? '-' : dateTime(data.startsAt).format('YYYY-MM-DD HH:mm:ss')}</>;
      },
      size: '180px',
    },
  ];
}

const getStyles = (theme: GrafanaTheme2) => ({
  table: css`
    max-width: ${theme.breakpoints.values.lg}px;
  `,
  moreMatches: css`
    margin-top: ${theme.spacing(1)};
  `,
  title: css`
    display: flex;
    align-items: center;
  `,
  badge: css`
    margin-left: ${theme.spacing(1)};
  `,
  stateColumn: css`
    display: flex;
    align-items: center;
  `,
});
