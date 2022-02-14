import React from "react";
import { Helmet } from "react-helmet";
import { FormProvider, useForm } from "react-hook-form";
import { Button } from "../../components/button/button";
import { Input } from "../../components/input/input";
import { Toggle } from "../../components/toggle/toggle";
import dayjs from "dayjs";
import { DevTool } from "@hookform/devtools";
import * as yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import "yup-phone";

export interface Props {}

const schema = yup.object({
  youtastePhone: yup
    .string()
    .required("Required")
    .phone("de", true, "Invalid phone number"),
  youtastePassword: yup.string().required("Required"),
  lieferandoUsername: yup.string().required("Required"),
  lieferandoPassword: yup.string().required("Required"),
});

/**
 * An AdminNewView React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const AdminNewView: React.FC<Props> = () => {
  const methods = useForm({
    defaultValues: {
      orderDatetime: dayjs().add(30, "minutes").format("YYYY-MM-DDTHH:mm:ss"),
      checkOpen: false,
    },
    resolver: yupResolver(schema),
  });

  const onSubmit = (data: any) => {};

  return (
    <FormProvider {...methods}>
      <div className="w-full h-full flex items-center justify-center">
        <Helmet>
          <title>Setup | TastyFood</title>
        </Helmet>
        <div className="background-card">
          <h1 className="background-card-title">Setup</h1>
          <form
            className="flex flex-col gap-y-2 mt-8"
            noValidate
            onSubmit={methods.handleSubmit(onSubmit)}
          >
            <div>
              <h2 className="text-lg font-medium mb-2">Youtaste Login</h2>
              <div className="flex gap-x-4">
                <Input
                  type="tel"
                  required
                  className="bg-transparent"
                  label="Phone"
                  name="youtastePhone"
                />
                <Input
                  required
                  type={"password"}
                  label="Password"
                  name="youtastePassword"
                />
              </div>
            </div>
            <div>
              <h2 className="text-lg font-medium mt-2 mb-2">
                Lieferando Login
              </h2>
              <div className="flex gap-x-4">
                <Input required label="Username" name="lieferandoUsername" />
                <Input
                  required
                  type={"password"}
                  label="Password"
                  name="lieferandoPassword"
                />
              </div>
            </div>
            <div className="relative right-10 w-112 mt-4  border-gray-100 border" />
            <div>
              <h2 className="text-lg font-medium mt-2 mb-2">Other Settings</h2>

              <Input
                required
                type={"datetime-local"}
                label="Order Datetime"
                name="orderDatetime"
              />
              <Toggle className="mt-2" name="checkOpen">
                Use Open Time of Restaurant (if available)
              </Toggle>

              <Button className="mt-4 ">Submit</Button>
            </div>
          </form>
        </div>
        <DevTool placement="top-right" control={methods.control} />
      </div>
    </FormProvider>
  );
};
