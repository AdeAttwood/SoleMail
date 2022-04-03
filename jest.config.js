/**
 * @copyright 2022 Practically.io
 */
module.exports = {
    preset: 'ts-jest',
    moduleNameMapper: {
        '\\.(css|scss)$': '<rootDir>/styleMock.js',
        '^@app/(.*)': '<rootDir>/frontend/$1',
    },
    coverageDirectory: '<rootDir>/runtime/jest-coverage',
    collectCoverageFrom: [
        '<rootDir>/src/**/*.{js,jsx,ts,tsx}',
        '!**/node_modules/**',
        '!**/vendor/**',
    ],
};
