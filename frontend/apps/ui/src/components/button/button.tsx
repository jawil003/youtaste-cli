import React from "react";
import { useFormContext } from "react-hook-form";

export interface Props extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  color?: "blue" | "red" | "green" | "gray" | "white";
  variant?: "primary" | "secondary";
}

//TODO: Check the way the bg color is generated (not good working with tailwind)

/**
 * An Button React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Button: React.FC<Props> = ({
  children,
  color,
  variant,
  className,
  ...props
}) => {
  if (props?.type === "submit") {
    const {
      formState: { errors },
      // eslint-disable-next-line react-hooks/rules-of-hooks
    } = useFormContext();

    const isDisabled = Object.values(errors).length > 0;

    if (variant === "primary")
      return (
        <button
          disabled={isDisabled}
          type="submit"
          className={
            `py-2 px-4  ${
              isDisabled
                ? "bg-gray-200"
                : `bg-${color}-600 hover:bg-${color}-700 focus:ring-${color}-500 focus:ring-offset-${color}-200`
            } text-white w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2  rounded-lg ` +
            className
          }
          {...props}
        >
          {children}
        </button>
      );
    else
      return (
        <button
          disabled={isDisabled}
          type="submit"
          className={
            `py-2 px-4  ${
              isDisabled
                ? "border-gray-200 "
                : `border-${color}-600 hover:border-${color}-700 focus:ring-${color}-500 focus:ring-offset-${color}-200`
            } text-black w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2  rounded-lg ` +
            className
          }
          {...props}
        >
          {children}
        </button>
      );
  }

  if (variant === "primary")
    return (
      <button
        type="submit"
        className={
          `py-2 px-4  ${`bg-${color}-600 hover:bg-${color}-700 focus:ring-${color}-500 focus:ring-offset-${color}-200`} text-white w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2  rounded-lg ` +
          className
        }
        {...props}
      >
        {children}
      </button>
    );
  else
    return (
      <button
        type="submit"
        className={
          `py-2 px-4  ${`border-${color}-600 hover:border-${color}-700 focus:ring-${color}-500 focus:ring-offset-${color}-200`} text-black w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2  rounded-lg ` +
          className
        }
        {...props}
      >
        {children}
      </button>
    );
};
Button.defaultProps = { color: "blue", variant: "primary" };
