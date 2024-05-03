import { css } from '@emotion/css';
import React from 'react';
import { useFormContext, useFieldArray, Controller } from 'react-hook-form';

import { GrafanaTheme2 } from '@grafana/data';
import { Button, Field, Input, IconButton, useStyles2, Select } from '@grafana/ui';
import { MatcherOperator } from 'app/plugins/datasource/alertmanager/types';

import { SilenceFormFields } from '../../types/silence-form';
import { matcherFieldOptions } from '../../utils/alertmanager';

interface Props {
  className?: string;
  required: boolean;
}

const MatchersField = ({ className, required }: Props) => {
  const styles = useStyles2(getStyles);
  const formApi = useFormContext<SilenceFormFields>();
  const {
    control,
    register,
    formState: { errors },
  } = formApi;

  const { fields: matchers = [], append, remove } = useFieldArray<SilenceFormFields>({ name: 'matchers' });

  return (
    <div className={className}>
      <Field label="Matching labels" required={required}>
        <div>
          <div className={styles.matchers}>
            {matchers.map((matcher, index) => {
              return (
                <div className={styles.row} key={`${matcher.id}`} data-testid="matcher">
                  <Field
                    label="Label"
                    invalid={!!errors?.matchers?.[index]?.name}
                    error={errors?.matchers?.[index]?.name?.message}
                  >
                    <Input
                      {...register(`matchers.${index}.name` as const, {
                        required: { value: required, message: 'Required.' },
                      })}
                      defaultValue={matcher.name}
                      placeholder="label"
                    />
                  </Field>
                  <Field label="Operator">
                    <Controller
                      control={control}
                      render={({ field: { onChange, ref, ...field } }) => (
                        <Select
                          {...field}
                          onChange={(value) => onChange(value.value)}
                          className={styles.matcherOptions}
                          options={matcherFieldOptions}
                          aria-label="operator"
                        />
                      )}
                      defaultValue={matcher.operator || matcherFieldOptions[0].value}
                      name={`matchers.${index}.operator`}
                      rules={{ required: { value: required, message: 'Required.' } }}
                    />
                  </Field>
                  <Field
                    label="Value"
                    invalid={!!errors?.matchers?.[index]?.value}
                    error={errors?.matchers?.[index]?.value?.message}
                  >
                    <Input
                      {...register(`matchers.${index}.value` as const, {
                        required: { value: required, message: 'Required.' },
                      })}
                      defaultValue={matcher.value}
                      placeholder="value"
                    />
                  </Field>
                  {(matchers.length > 1 || !required) && (
                    <IconButton
                      aria-label="Remove matcher"
                      className={styles.removeButton}
                      name="trash-alt"
                      onClick={() => remove(index)}
                    >
                      Remove
                    </IconButton>
                  )}
                </div>
              );
            })}
          </div>
          <Button
            tooltip="Refine which alert instances are silenced by selecting label matchers"
            type="button"
            icon="plus"
            variant="secondary"
            onClick={() => {
              const newMatcher = { name: '', value: '', operator: MatcherOperator.equal };
              append(newMatcher);
            }}
          >
            Add matcher
          </Button>
        </div>
      </Field>
    </div>
  );
};

const getStyles = (theme: GrafanaTheme2) => {
  return {
    row: css({
      display: 'flex',
      alignItems: 'flex-start',
      flexDirection: 'row',
      backgroundColor: theme.colors.background.secondary,
      padding: `${theme.spacing(1)} ${theme.spacing(1)} 0 ${theme.spacing(1)}`,
      gap: theme.spacing(1),
    }),
    removeButton: css({
      marginLeft: theme.spacing(1),
      marginTop: theme.spacing(3),
    }),
    matcherOptions: css({
      minWidth: '140px',
    }),
    matchers: css({
      maxWidth: theme.breakpoints.values.sm,
      margin: `${theme.spacing(1)} 0`,
    }),
  };
};

export default MatchersField;
