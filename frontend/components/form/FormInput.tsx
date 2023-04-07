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
     * Is the FormInput required?
     */
    required?: boolean;
}

/**
 * UI component for text inputs
 */
export const FormInput = ({
    id,
    type,
    name,
    label,
    required = true,
    ...props
}: FormInputProps) => {
    return (<>
        <label htmlFor="name">{label}</label>
        <input
            id={id}
            className="p-2 rounded-md bg-white mb-2"
            {...props}
        />
    </>
    );
}