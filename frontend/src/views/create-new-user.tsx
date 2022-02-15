import React, { useEffect, useState } from "react";
import { FormProvider, useForm } from "react-hook-form";
import { Button } from "../components/button/button";
import { Input } from "../components/input/input";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import { useNavigate } from "react-router-dom";
import { Helmet } from "react-helmet";
import { useStore } from "../store/store";
import UserService from "../services/user.service";
import { useQueryClient } from "react-query";
import { Queries } from "../enums/queries.enum";
import { useTranslation } from "react-i18next";
import { useRedirector } from "../hooks/useRedirector";

export interface Props {}

const schema = yup.object().shape({
  firstname: yup.string().required(),
  lastname: yup.string().required(),
});

interface FormData {
  firstname: string;
  lastname: string;
}

/**
 * An CreateNewUserView React Component.
 * @author
 * @version 0.1
 */
export const CreateNewUserView: React.FC<Props> = () => {
  const [loadingButtonEnabled, setLoadingButtonEnabled] = useState(false);
  const { setUser } = useStore();
  const route = useRedirector();

  const { t } = useTranslation("create-new-user");

  useEffect(() => {
    const firstname = localStorage.getItem("firstname");

    setLoadingButtonEnabled(Boolean(firstname));
  }, [setLoadingButtonEnabled]);

  const queryClient = useQueryClient();

  const navigate = useNavigate();
  const methods = useForm<FormData>({
    defaultValues: { firstname: "", lastname: "" },
    resolver: yupResolver(schema),
  });

  const onSubmit = async (value: FormData) => {
    localStorage.setItem("firstname", value.firstname);
    localStorage.setItem("lastname", value.lastname);

    const userService = new UserService();

    await userService.create(value.firstname, value.lastname);

    setUser(value);
    await queryClient.invalidateQueries(Queries.USER);

    methods.reset();
    navigate(route);
  };

  return (
    <div className="w-full h-full flex items-center justify-center">
      <Helmet>
        <title>{t("headline")} | TastyFood</title>
      </Helmet>
      <div className="background-card flex flex-col">
        <div className="background-card-title">{t("headline")}</div>
        <div className="mt-8">
          <FormProvider {...methods}>
            <form noValidate onSubmit={methods.handleSubmit(onSubmit)}>
              <Input
                required
                name="firstname"
                label={t("firstname")}
                placeholder={t("firstnamePlaceholder")}
              />
              <div className="h-2" />
              <Input
                required
                name="lastname"
                label={t("lastname")}
                placeholder={t("lastnamePlaceholder")}
              />
              <div className="flex w-full mt-8 flex-col">
                <Button className="mb-2">{t("submit")}</Button>
                {loadingButtonEnabled && (
                  <Button
                    onClick={() => {
                      const firstname = localStorage.getItem("firstname");
                      const lastname = localStorage.getItem("lastname");

                      if (firstname && lastname) {
                        methods.setValue("firstname", firstname);
                        methods.setValue("lastname", lastname);
                      }
                    }}
                    type="button"
                    variant="secondary"
                  >
                    {t("load")}
                  </Button>
                )}
              </div>
            </form>
          </FormProvider>
        </div>
      </div>
    </div>
  );
};
