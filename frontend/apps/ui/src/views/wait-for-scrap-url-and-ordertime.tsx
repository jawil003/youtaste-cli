import React from "react";
import { Helmet } from "react-helmet";
import { useTranslation } from "react-i18next";
import { ReactComponent as PizzaDelivery } from "../assets/Pizza sharing-rafiki.svg";
export interface Props {}

/**
 * An WaitForScrapUrlAndOrdertimeView React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const WaitForScrapUrlAndOrdertimeView: React.FC<Props> = () => {
  const { t } = useTranslation("wait-for-scrap-url-and-ordertime");

  return (
    <div className="bg-white w-full h-full px-12 mt-6 flex flex-row justify-between items-center relative">
      <Helmet>
        <title>{t("headline")} | TastyFood</title>
      </Helmet>
      <div className="flex-1">
        <div className="flex flex-1 m-auto max-w-lg  justify-center items-start flex-col mb-16">
          <h1 className="font-thin max-w-lg text-6xl text-gray-800">
            {t("waitToChooseRestaurant")}
          </h1>
          <p className="font-light text-xl text-gray-800 mt-4">
            {t("willInformYouSoon")}
          </p>
        </div>
      </div>
      <div className="block flex-1 mx-auto mt-6 md:mt-0 relative">
        <PizzaDelivery width={"100%"} />
      </div>
    </div>
  );
};
