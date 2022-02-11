import React from "react";
import { BsFillTrashFill, BsPencilFill } from "react-icons/bs";
export interface Props {
  headline: string;
  items: {
    headline: string;
    description: string;
    onEditClick?: React.MouseEventHandler<HTMLButtonElement>;
    onDeleteClick?: React.MouseEventHandler<HTMLButtonElement>;
    onClick?: React.MouseEventHandler<HTMLButtonElement>;
    size: number;
  }[];
}

/**
 * An OrderList React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const OrderList: React.FC<Props> = ({ headline, children, items }) => {
  return (
    <div className="flex flex-col max-w-md p-8 bg-white rounded-lg shadow max-h-3/4">
      <div className="background-card-title">{headline}</div>

      <div className="container flex flex-col mx-auto items-center overflow-x-hidden overflow-y-scroll w-80">
        <ul className="flex flex-col w-full gap-2">
          {items.map(
            ({
              description,
              headline,
              onClick,
              size,
              onDeleteClick,
              onEditClick,
            }) => (
              <li className="border-gray-400 flex flex-row">
                <div className="shadow border select-none cursor-pointer bg-white rounded-md flex flex-1 items-center p-4">
                  <div className="flex-1 pl-1 md:mr-16">
                    <div className="font-medium">{headline}</div>
                    <div className="text-gray-600 text-sm">{description}</div>
                  </div>
                  <div className="text-gray-600 text-xs">{size}x</div>
                  {onEditClick && (
                    <button onClick={onEditClick}>
                      <BsPencilFill className="ml-4" />
                    </button>
                  )}
                  {onDeleteClick && (
                    <button onClick={onDeleteClick}>
                      <BsFillTrashFill className="ml-2" />
                    </button>
                  )}
                  {onClick && (
                    <button
                      className="w-24 text-right flex justify-end"
                      onClick={onClick}
                    >
                      <svg
                        width="12"
                        fill="currentColor"
                        height="12"
                        className="hover:text-gray-800 text-gray-500"
                        viewBox="0 0 1792 1792"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path d="M1363 877l-742 742q-19 19-45 19t-45-19l-166-166q-19-19-19-45t19-45l531-531-531-531q-19-19-19-45t19-45l166-166q19-19 45-19t45 19l742 742q19 19 19 45t-19 45z"></path>
                      </svg>
                    </button>
                  )}
                </div>
              </li>
            )
          )}
        </ul>
      </div>
      {children}
    </div>
  );
};
