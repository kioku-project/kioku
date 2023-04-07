import React from 'react';

interface FormButtonProps {
    /**
     * unique identifier
     */
    id: string;
    /**
     * Button contents
     */
    value: string;
    /**
     * optional click handler
     */
    onClick?: () => void;
}

/**
 * UI component for submitting forms
 */
export const FormButton = ({
    id,
    value,
    ...props
}: FormButtonProps) => {
    return (
        <input
            type="submit"
            id={id}
            value={value}
            className="p-2 rounded-md bg-white hover:cursor-pointer transition ease-in-out delay-100 hover:-translate-y-0.5 hover:scale-105 hover:bg-gray-100 duration-200"
            {...props}
        />
    );
}