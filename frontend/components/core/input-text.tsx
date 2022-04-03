import React from 'react';
import {InputHTMLAttributes} from 'react';

export interface InputTextProps extends InputHTMLAttributes<HTMLInputElement> {
    /**
     * All inputs must a unique ID to link labels and hints to the input
     */
    id: string;
    /**
     * The label that will describe the input
     */
    label?: string;
    /**
     * The type of input this is. The Default value is 'text'
     */
    type?: 'text' | 'email' | 'number' | 'password';
}

export const InputText: React.FC<InputTextProps> = (props) => {
    const {label, ...rest} = props;
    return (
        <div className="w-full">
            {label && (
                <label className="label" htmlFor={rest.id}>
                    <span className="label-text">{label}</span>
                </label>
            )}
            <input
                {...rest}
                className="w-full px-4 py-1 text-gray-800 rounded-full focus:outline-none shadow-lg border"
            />
        </div>
    );
};

export default InputText;
