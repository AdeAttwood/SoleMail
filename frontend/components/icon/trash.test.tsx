import React from 'react';
import {render, cleanup} from '@testing-library/react';
import '@testing-library/jest-dom';

import Icon from './trash';

afterEach(cleanup);

test('Trash icon renders', () => {
    const {getByRole, container} = render(<Icon role={'button'} />);
    expect(getByRole('button')).toBeInTheDocument();
});
