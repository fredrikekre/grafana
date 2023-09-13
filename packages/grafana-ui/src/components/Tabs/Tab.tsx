import { css, cx, keyframes } from '@emotion/css';
import React, { HTMLProps } from 'react';

import { GrafanaTheme2, NavModelItem } from '@grafana/data';
import { selectors } from '@grafana/e2e-selectors';

import { useStyles2 } from '../../themes';
import { getFocusStyles } from '../../themes/mixins';
import { IconName } from '../../types';
import { Icon } from '../Icon/Icon';

import { Counter } from './Counter';

export interface TabProps extends HTMLProps<HTMLAnchorElement> {
  label: string;
  active?: boolean;
  /** When provided, it is possible to use the tab as a hyperlink. Use in cases where the tabs update location. */
  href?: string;
  icon?: IconName;
  onChangeTab?: (event?: React.MouseEvent<HTMLAnchorElement>) => void;
  /** A number rendered next to the text. Usually used to display the number of items in a tab's view. */
  counter?: number | null;
  /** Extra content, displayed after the tab label and counter */
  suffix?: NavModelItem['tabSuffix'];
}

export const Tab = React.forwardRef<HTMLAnchorElement, TabProps>(
  ({ label, active, icon, onChangeTab, counter, suffix: Suffix, className, href, ...otherProps }, ref) => {
    const tabsStyles = useStyles2(getStyles);
    const content = () => (
      <>
        {icon && <Icon name={icon} />}
        {label}
        {typeof counter === 'number' && <Counter value={counter} />}
        {Suffix && <Suffix className={tabsStyles.suffix} />}
      </>
    );

    const containerClass = cx(tabsStyles.item, active ? tabsStyles.itemActive : null);
    const linkClass = cx(tabsStyles.link, active ? tabsStyles.activeStyle : tabsStyles.notActive, className);

    return (
      <div className={containerClass}>
        <a
          // in case there is no href '#' is set in order to maintain a11y
          href={href ? href : '#'}
          className={linkClass}
          {...otherProps}
          onClick={onChangeTab}
          aria-label={otherProps['aria-label'] || selectors.components.Tab.title(label)}
          role="tab"
          aria-selected={active}
          ref={ref}
        >
          {content()}
        </a>
      </div>
    );
  }
);

Tab.displayName = 'Tab';

const getStyles = (theme: GrafanaTheme2) => {

  // const gradientAnim = keyframes`
  //   0% {
  //     background-position: 0% 50%;
  //   }
  //   50% {
  //     background-position: 100% 50%;
  //   }
  //   100% {
  //     background-position: 0% 50%;
  //   }
  // `;


  return {
    item: css({
      listStyle: 'none',
      position: 'relative',
      whiteSpace: 'nowrap',
      // backgroundColor: `${theme.colors.background.secondary}`,
      marginBottom: '0px',

      '&::before': {
        display: 'block',
        content: '" "',
        height: '4px',
        width: '100%',
        borderRadius: `${theme.shape.radius.default} ${theme.shape.radius.default} 0 0`,
        position: 'absolute',
        top: 0,
        transition: 'all 0.5s ease',
      },

      '&:hover': {
        a: {
          backgroundColor: theme.colors.background.primary,
          borderBottom: `1px solid ${theme.components.panel.borderColor}`,
        }
      },

      'a:hover, &:hover, &:focus': {
        color: theme.colors.text.primary,

        '&::before': {
          borderRadius: '3px 3px 0 0',
          backgroundColor: theme.colors.action.hover,
        },
      },
    }),
    itemActive: css({
      backgroundColor: `${theme.colors.background.primary}`,

      '&:hover': {

        '&::before': {
          boxShadow: '-2px 0px 2px #FF8833, 2px 0px 2px #F55F3E',
        },

        a: {
          backgroundColor: theme.colors.background.primary,
          borderBottom: 0,
        }
      },

      '&::before': {
        backgroundImage: theme.colors.gradients.brandHorizontal,
      },
    }),
    link: css({
      color: theme.colors.text.secondary,
      padding: theme.spacing(1.5, 2, 1),
      display: 'block',
      height: 'calc(100%)',
      // borderBottom: `1px solid ${theme.components.panel.borderColor}`,

      svg: {
        marginRight: theme.spacing(1),
      },

      '&:focus-visible': getFocusStyles(theme),

      
    }),
    notActive: css({
      
    }),
    activeStyle: css({
      label: 'activeTabStyle',
      color: theme.colors.text.primary,
      overflow: 'hidden',
      border: 0,
      borderLeft: `1px solid ${theme.components.panel.borderColor}`,
      borderRight: `1px solid ${theme.components.panel.borderColor}`,

      a: {
        color: theme.colors.text.primary,
      },
    }),
    suffix: css({
      marginLeft: theme.spacing(1),
    }),
  };
};
