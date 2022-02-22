import React from "react";
import { Helmet } from "react-helmet";
import { useTranslation } from "react-i18next";
import { ReactComponent as BarbecueAmico } from "../assets/Barbecue-amico.svg";

export interface Props {}

/**
 * An OrderInProgressView React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const OrderInProgressView: React.FC<Props> = () => {
  const { t } = useTranslation("order-in-progress");

  return (
    <div className="bg-white w-full h-full px-12 mt-6 flex flex-row justify-between items-center relative">
      <Helmet>
        <title>{t("headline")} | TastyFood</title>
      </Helmet>
      <div className="flex-1">
        <div className="flex flex-1 m-auto max-w-lg  justify-center items-start flex-col mb-16">
          <h1 className="font-thin max-w-lg text-6xl text-gray-800">
            {t("orderInProgress")}
          </h1>
          <p className="font-light text-xl text-gray-800 mt-4">
            {t("orderIsSoonDone")}
          </p>
        </div>
      </div>
      <div className="block flex-1 mx-auto mt-6 md:mt-0 relative">
        <BarbecueAmico width={"100%"} />
      </div>
    </div>
  );
};
