import React from 'react';
import classes from '@app/lib/classes';

export interface ButtonProps
    extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    tooltip?: string;
}

export const Button: React.FC<ButtonProps> = ({children, tooltip, ...rest}) => {
    return (
        <button
            {...rest}
            className={classes([
                'btn btn-primary no-animation',
                tooltip ? 'tooltip' : false,
            ])}
            data-tip={tooltip}
        >
            {children}
        </button>
    );
};

export default Button;
