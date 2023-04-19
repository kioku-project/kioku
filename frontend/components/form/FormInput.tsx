import React from 'react';

interface FormInputProps {
    /**
     * unique identifier
     */
    id: string;
    /**
     * FormInput type
     */
    type: string;
    /**
    * FormInput name
    */
    name: string;
    /**
     * optional FormInput label
     */
    label?: string;
    /**
     * FormInput value
     */
    value?: string;
    /**
     * Is the FormInput required?
     */
    required?: boolean;
    /**
     * Additional classes
     */
    className?: string;
}

/**
 * UI component for text inputs
 */
export const FormInput = ({
    name,
    label,
    required = true,
    className,
    ...props
}: FormInputProps) => {
    return (<>
        <label htmlFor={name}>{label}</label>
        <input
            name={name}
            className={`p-2 rounded-md bg-white mb-2 ${className || ''}`}
            required={required}
            {...props}
        />
    </>
    );
}