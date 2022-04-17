import * as React from 'react';
import {SVGProps} from 'react';

const ExclamationMark = (props: SVGProps<SVGSVGElement>) => (
    <svg
        width={40}
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 64 64"
        xmlSpace="preserve"
        {...props}
    >
        <path d="M36.989 42.439H27.01L23 2h18z" fill="currentColor" />
        <ellipse
            cx={31.999}
            cy={54.354}
            rx={7.663}
            ry={7.646}
            fill="currentColor"
        />
    </svg>
);

export default ExclamationMark;
