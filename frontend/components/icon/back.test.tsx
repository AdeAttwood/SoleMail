import React from 'react';
import {render, cleanup} from '@testing-library/react';
import '@testing-library/jest-dom';

import Icon from './back';

afterEach(cleanup);

test('Back icon renders', () => {
    const {getByRole, container} = render(<Icon role={'button'} />);
    expect(getByRole('button')).toBeInTheDocument();
});
