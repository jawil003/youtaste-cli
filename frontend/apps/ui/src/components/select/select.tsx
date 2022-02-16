import React, { useState } from "react";
import { useFormContext, useWatch } from "react-hook-form";

export interface Props {
  label?: string;
  name: string;
  placeholder?: string;
  required?: boolean;
  className?: string;
  options: { value: string; label: string }[];
}

/**
 * An Select React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Select: React.FC<Props> = ({
  label,
  required,
  className,
  placeholder,
  options,
  name,
}) => {
  const [open, setOpen] = useState(false);
  const {
    setValue,
    formState: { errors },
  } = useFormContext();

  const value = useWatch({ name }) as string | null | undefined;

  const isError = errors?.[name]
    ? Object?.values?.(errors?.[name])?.length > 0
    : false;

  return (
    <div className={className}>
      <div className=" relative">
        {" "}
        {label && (
          <>
            <label className="text-gray-700">{label}</label>
            {required && (
              <span
                className={
                  isError ? "text-red-500 required-dot" : "text-gray-700"
                }
              >
                *
              </span>
            )}
          </>
        )}
        <div>
          <div className="relative">
            <button
              onClick={() => setOpen(!open)}
              type="button"
              className="relative w-full bg-white rounded-md border border-gray-300 py-2 px-4 pr-10 text-left cursor-default focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
            >
              <span className="flex items-center">
                <span
                  className={`block truncate ${
                    !value && placeholder ? "text-gray-400" : "text-gray-700"
                  } text-base`}
                >
                  {options.find(({ value: valueOptions, label }) => {
                    if (value === valueOptions) return true;
                    return false;
                  })?.label ??
                    placeholder ??
                    "Please select..."}
                </span>
              </span>
              <span className="absolute inset-y-0 right-0 flex items-center pr-2 pointer-events-none">
                <svg
                  className="h-5 w-5 text-gray-400"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                  aria-hidden="true"
                >
                  <path
                    fill-rule="evenodd"
                    d="M10 3a1 1 0 01.707.293l3 3a1 1 0 01-1.414 1.414L10 5.414 7.707 7.707a1 1 0 01-1.414-1.414l3-3A1 1 0 0110 3zm-3.707 9.293a1 1 0 011.414 0L10 14.586l2.293-2.293a1 1 0 011.414 1.414l-3 3a1 1 0 01-1.414 0l-3-3a1 1 0 010-1.414z"
                    clip-rule="evenodd"
                  ></path>
                </svg>
              </span>
            </button>
            {open && (
              <div className="absolute mt-1 w-full z-10 rounded-md bg-white shadow-lg">
                <ul
                  tabIndex={-1}
                  role="listbox"
                  aria-labelledby="listbox-label"
                  aria-activedescendant="listbox-item-3"
                  className="max-h-56 rounded-md py-1 text-base ring-1 ring-black ring-opacity-5 overflow-auto focus:outline-none sm:text-sm"
                >
                  {options.map(({ label, value: valueOption }) => (
                    <li
                      onClick={() => {
                        setValue(name, valueOption);
                        setOpen(false);
                      }}
                      id="listbox-item-0"
                      // eslint-disable-next-line jsx-a11y/role-has-required-aria-props
                      role="option"
                      className="text-gray-900 cursor-default hover:bg-blue-500 hover:text-white select-none relative py-2 pl-3 pr-9"
                    >
                      <div className="flex items-center">
                        <span className="ml-3 block font-normal truncate">
                          {label}
                        </span>
                      </div>
                      {value === valueOption && (
                        <span className="absolute inset-y-0 right-0 flex items-center pr-4">
                          <svg
                            className="h-5 w-5"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 20 20"
                            fill="currentColor"
                            aria-hidden="true"
                          >
                            <path
                              fill-rule="evenodd"
                              d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                              clip-rule="evenodd"
                            ></path>
                          </svg>
                        </span>
                      )}
                    </li>
                  ))}
                </ul>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};
