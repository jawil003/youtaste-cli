import React from "react";
import { useFormContext } from "react-hook-form";
import classNames from "classnames";
export interface Props extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  color?: "blue" | "red" | "green" | "white";
}

/**
 * An Button React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Button: React.FC<Props> = ({
  children,
  color,

  className,
  ...props
}) => {
  let isDisabled = false;

  const calcClassNames = classNames({
    "py-2 px-4 w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2 rounded-lg":
      true,
    "bg-gray-200": isDisabled,
    "text-white": color !== "white",
    "text-blue": color === "white",
    "bg-blue-600 hover:bg-blue-700 focus:ring-blue-500 focus:ring-offset-blue-200":
      color === "blue" && !isDisabled,
    "bg-green-600 hover:bg-green-700 focus:ring-green-500 focus:ring-offset-green-200":
      color === "green" && !isDisabled,
    "bg-red-600 hover:bg-red-700 focus:ring-red-500 focus:ring-offset-red-200":
      color === "red" && !isDisabled,
    "border-blue-600 hover:border-blue-700 focus:ring-blue-500 focus:ring-offset-blue-200":
      color === "white" && !isDisabled,
  });

  if (props?.type === "submit") {
    const {
      formState: { errors },
      // eslint-disable-next-line react-hooks/rules-of-hooks
    } = useFormContext();

    isDisabled = Object.values(errors).length > 0;

    return (
      <button
        disabled={isDisabled}
        type="submit"
        className={calcClassNames + (className ? ` ${className}` : "")}
        {...props}
      >
        {children}
      </button>
    );
  }

  return (
    <button
      type="submit"
      className={calcClassNames + (className ? ` ${className}` : "")}
      {...props}
    >
      {children}
    </button>
  );
};
Button.defaultProps = { color: "blue" };
