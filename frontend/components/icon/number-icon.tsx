import * as React from 'react';
import {SVGProps} from 'react';

export interface NumberIconProps extends SVGProps<SVGSVGElement> {
    /**
     * The number to display in the icon
     */
    number: number;
    /**
     * If the component is in a spinning state. The default value is false.
     */
    spinning?: boolean;
}

const NumberIcon = ({spinning, number, ...rest}: NumberIconProps) => (
    <svg
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 100 59"
        xmlSpace="preserve"
        {...rest}
    >
        <path
            d="M0 7.393V49.995h69.618V7.393H0Zm2.785 2.582h2.284l29.74 23.48 29.74-23.48h2.284v1.594l-20.907 16.5 20.907 18.255v1.089h-2.806l-20.233-17.65-8.985 7.08-8.985-7.08L5.59 47.413H2.785v-1.09L23.692 28.07l-20.907-16.5V9.975Z"
            fill="#62D2A2"
        />
        <circle
            cx={70.125}
            cy={29.201}
            r={28.408}
            fill="#62D2A2"
            stroke="#3AC78B"
        />
        {spinning ? (
            <path
                style={{transformOrigin: '71% 49%'}}
                className="animate-spin"
                d="M62.438 37.644a11.768 11.768 0 0 1-3.027-4.33c-1.73-4.242-.92-9.287 2.457-12.766l3.116 3.116c.172.171.48.048.522-.206l1.702-11.14a.277.277 0 0 0-.316-.315l-11.146 1.702c-.254.041-.377.35-.206.522l3.123 3.123c-5.154 5.27-6.033 13.177-2.65 19.354a16.472 16.472 0 0 0 4.215 4.99c3.047 2.422 6.794 3.582 10.514 3.506l.7-4.577c-3.205.274-6.5-.714-9.004-2.979ZM83.165 40.396c5.154-5.27 6.032-13.177 2.649-19.354a16.472 16.472 0 0 0-4.214-4.99c-3.047-2.422-6.795-3.582-10.515-3.507l-.7 4.578c3.199-.274 6.5.714 9.005 2.979a11.767 11.767 0 0 1 3.026 4.33c1.73 4.242.92 9.286-2.456 12.766l-3.116-3.116c-.172-.172-.48-.048-.522.206L74.62 45.427a.277.277 0 0 0 .316.316L86.08 44.04c.254-.042.371-.35.206-.522l-3.122-3.123Z"
                fill="#fff"
            />
        ) : (
            <text
                x={55}
                y={number >= 100 ? 35 : 38}
                fill="#fff"
                fontWeight="bold"
                fontSize={number >= 100 ? 20 : 30}
            >
                {number}
            </text>
        )}
    </svg>
);

export default NumberIcon;
