/**
 * Days of the week in words. This array is in the order of the `Date.getDay`
 * function
 */
const DAYS = [
    'Sunday',
    'Monday',
    'Tuseday',
    'Wednesday',
    'Thursday',
    'Friday',
    'Saturday',
];

/**
 * Formats the date
 */
const formatDateString = (date: Date): string => {
    return [date.getDate(), date.getMonth() + 1, date.getFullYear()].join('/');
};

const formatTimeString = (date: Date): string => {
    let hours: string | number = date.getHours();
    let minutes: string | number = date.getMinutes();

    if (hours < 10) {
        hours = '0' + hours.toString();
    }

    if (minutes < 10) {
        minutes = '0' + minutes.toString();
    }
    return `${hours}:${minutes}`;
};

export const formatFullString = (date: string): string => {
    const date_time = new Date(date);
    return `${formatDateString(date_time)} ${formatTimeString(date_time)}`;
};

export const formatDate = (date: string): string => {
    const test_date = new Date();
    const date_time = new Date(date);

    if (formatDateString(test_date) == formatDateString(date_time)) {
        return `Today ${formatTimeString(date_time)}`;
    }

    test_date.setDate(test_date.getDate() - 1);
    if (formatDateString(test_date) == formatDateString(date_time)) {
        return `Yesterday ${formatTimeString(date_time)}`;
    }

    test_date.setDate(test_date.getDate() - 4);
    if (test_date < date_time) {
        const day = DAYS[date_time.getDay()];
        return `${day} ${formatTimeString(date_time)}`;
    }

    return `${formatDateString(date_time)} ${formatTimeString(date_time)}`;
};

export default formatDate;
