import React, { useEffect, useState } from "react";
import { FormProvider, useForm } from "react-hook-form";
import { Button } from "../components/button/button";
import { Input } from "../components/input/input";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import { useNavigate } from "react-router-dom";
import { Routes } from "../enums/routes.enum";
import { Helmet } from "react-helmet";

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

  useEffect(() => {
    const firstname = localStorage.getItem("firstname");

    setLoadingButtonEnabled(Boolean(firstname));
  }, [setLoadingButtonEnabled]);

  const navigate = useNavigate();
  const methods = useForm<FormData>({
    defaultValues: { firstname: "", lastname: "" },
    resolver: yupResolver(schema),
  });

  const onSubmit = (value: FormData) => {
    localStorage.setItem("firstname", value.firstname);
    localStorage.setItem("lastname", value.lastname);
    methods.reset();
    navigate(Routes.NEW_ORDER);
  };

  return (
    <div className="w-screen h-screen flex items-center justify-center">
      <Helmet>
        <title>New User | TastyFood</title>
      </Helmet>
      <div className="flex flex-col w-full max-w-md px-4 py-8 bg-white rounded-lg shadow dark:bg-gray-800 sm:px-6 md:px-8 lg:px-10">
        <div className="self-center text-xl font-light text-gray-600 sm:text-2xl dark:text-white">
          Who are you?
        </div>
        <div className="mt-8">
          <FormProvider {...methods}>
            <form noValidate onSubmit={methods.handleSubmit(onSubmit)}>
              <Input
                required
                name="firstname"
                label="Firstname"
                placeholder="Maxine"
              />
              <div className="h-2" />
              <Input
                required
                name="lastname"
                label="Lastname"
                placeholder="Smith"
              />
              <div className="flex w-full mt-8 flex-col">
                <Button className="mb-2">Register</Button>
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
                    Load
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
