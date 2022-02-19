import React from "react";
import { CheckIcon, XIcon, DotsHorizontalIcon } from "@heroicons/react/solid";
import { useLocation } from "react-router-dom";
import { Routes } from "../../enums/routes.enum";
import { useTranslation } from "react-i18next";
import { useIsAdmin } from "../../hooks/isAdmin.hook";
export interface Props {
  stateFactory: {
    activeHref: string;
    states: { href: string; state: "DISABLED" | "AVAILABLE" | "DONE" }[];
  }[];
  items: {
    label: string;
    href: string;
  }[];
}

const resolveIconAndColor: (state?: "DISABLED" | "AVAILABLE" | "DONE") => {
  color: string;
  icon: React.ReactNode;
} = (state) => {
  switch (state) {
    case "DISABLED": {
      return {
        color: "bg-red-400",
        icon: <XIcon width={"0.8rem"} className="text-white" />,
      };
    }
    case "DONE": {
      return {
        color: "bg-green-400",
        icon: <CheckIcon width={"0.8rem"} className="text-white" />,
      };
    }
    default: {
      return {
        color: "bg-blue-400",
        icon: <DotsHorizontalIcon width={"0.8rem"} className="text-white" />,
      };
    }
  }
};

/**
 * An PositionTracker React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const PositionTracker: React.FC<Props> = ({ items, stateFactory }) => {
  const { pathname } = useLocation();

  const currentState = stateFactory.find((a) => a.activeHref === pathname);

  return (
    <ol className="ml-8 h-full w-40 mt-24 relative border-l border-gray-200">
      {items?.map(({ label, href }) => {
        const elementState = currentState?.states.find(
          (state) => state.href === href
        )?.state;

        const { color, icon } = resolveIconAndColor(elementState);

        return (
          <li className="mb-10 ml-6" key={href}>
            <span
              className={`flex absolute -left-3 justify-center items-center w-6 h-6 ${color} rounded-full ring-8 ring-white`}
            >
              {icon}
            </span>
            <h3 className="flex items-center mb-1 text-base font-normal text-gray-900">
              {label}
            </h3>
          </li>
        );
      })}
    </ol>
  );
};

export const PositionTrackerDefault: React.FC = () => {
  const { t } = useTranslation("position-tracker-default");
  const { data: isAdmin } = useIsAdmin();

  return (
    <PositionTracker
      stateFactory={[
        {
          activeHref: Routes.ADMIN_NEW,
          states: [
            {
              href: Routes.ADMIN_NEW,
              state: "AVAILABLE",
            },
            {
              href: Routes.NEW,
              state: "DISABLED",
            },
            { href: Routes.POLLS, state: "DISABLED" },
            { href: Routes.ORDER_CONFIRM, state: "DISABLED" },
            { href: Routes.ON_THE_WAY, state: "DISABLED" },
          ],
        },
        {
          activeHref: Routes.NEW,
          states: [
            {
              href: Routes.ADMIN_NEW,
              state: "DONE",
            },
            {
              href: Routes.NEW,
              state: "AVAILABLE",
            },
            { href: Routes.POLLS, state: "DISABLED" },
            { href: Routes.ORDER_CONFIRM, state: "DISABLED" },
            { href: Routes.ON_THE_WAY, state: "DISABLED" },
          ],
        },
        {
          activeHref: Routes.POLLS,
          states: [
            {
              href: Routes.ADMIN_NEW,
              state: "DONE",
            },
            {
              href: Routes.NEW,
              state: "DONE",
            },
            { href: Routes.POLLS, state: "AVAILABLE" },
            { href: Routes.ORDER_CONFIRM, state: "DISABLED" },
            { href: Routes.ON_THE_WAY, state: "DISABLED" },
          ],
        },
        {
          activeHref: Routes.ORDER_CONFIRM,
          states: [
            {
              href: Routes.ADMIN_NEW,
              state: "DONE",
            },
            {
              href: Routes.NEW,
              state: "DONE",
            },
            { href: Routes.POLLS, state: "DONE" },
            { href: Routes.ORDER_CONFIRM, state: "AVAILABLE" },
            { href: Routes.ON_THE_WAY, state: "DISABLED" },
          ],
        },
        {
          activeHref: Routes.ON_THE_WAY,
          states: [
            {
              href: Routes.ADMIN_NEW,
              state: "DONE",
            },
            {
              href: Routes.NEW,
              state: "DONE",
            },
            { href: Routes.POLLS, state: "DONE" },
            { href: Routes.ORDER_CONFIRM, state: "DONE" },
            { href: Routes.ON_THE_WAY, state: "AVAILABLE" },
          ],
        },
      ]}
      items={(() => {
        const base = [
          {
            href: Routes.NEW,
            label: t("comeIn"),
          },
          {
            href: Routes.POLLS,
            label: t("voteResteraunt"),
          },
          { href: Routes.ORDER_CONFIRM, label: t("myMeals") },
          {
            href: Routes.ON_THE_WAY,
            label: t("orderOnTheWay"),
          },
        ];

        const admin = isAdmin
          ? [
              {
                href: Routes.ADMIN_NEW,
                label: t("setup"),
              },
            ]
          : [];

        return [...admin, ...base];
      })()}
    />
  );
};
