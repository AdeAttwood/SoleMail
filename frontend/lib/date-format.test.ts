import {formatDate, formatFullString} from './date-format';

test('Format date full string', () => {
    const formatted = formatFullString('2022-01-01T10:10:00');
    expect(formatted).toBe('1/1/2022 10:10');
});

test('Format today', () => {
    const date = new Date();
    date.setHours(10);
    date.setMinutes(10);
    const formatted = formatDate(date.toUTCString());

    expect(formatted).toBe('Today 10:10');
});

test('Format yesterday', () => {
    const date = new Date();
    date.setDate(date.getDate() - 1);
    date.setHours(10);
    date.setMinutes(10);
    const formatted = formatDate(date.toUTCString());

    expect(formatted).toBe('Yesterday 10:10');
});

test('Format date from more than a week ago', () => {
    const formatted = formatDate('2022-01-01T10:10:00');
    expect(formatted).toBe('1/1/2022 10:10');
});
