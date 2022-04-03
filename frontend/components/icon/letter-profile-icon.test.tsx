import React from 'react';
import {render, cleanup} from '@testing-library/react';
import '@testing-library/jest-dom';

import Icon from './letter-profile-icon';

afterEach(cleanup);

test.each([
    ['Jane Doe', 'JD'],
    ['Jane', 'J'],
    ['Jane (@jane)', 'J'],
    ['(Cron Daemon)', 'CD'],
])('Test name of "%s" evaluates to "%s"', (name, letters) => {
    const {getByText} = render(<Icon name={name} />);
    expect(getByText(letters)).toBeInTheDocument();
});
