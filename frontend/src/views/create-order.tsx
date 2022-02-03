import React from "react";
import { Helmet } from "react-helmet";
import { FormProvider, useForm, useWatch } from "react-hook-form";
import { string } from "yup/lib/locale";
import { Logo } from "../assets/logo/logo";
import { Badge } from "../components/badge/badge";
import { Button } from "../components/button/button";
import { Input } from "../components/input/input";

export interface Props {}

interface FormData {
  mealName: string;
  variant: string;
  variants: string[];
}

/**
 * An CreateOrderView React Component.
 * @author
 * @version 0.1
 */
export const CreateOrderView: React.FC<Props> = () => {
  const methods = useForm<FormData>({
    defaultValues: { mealName: "", variant: "", variants: [] },
  });

  const variant = useWatch({ name: "variant", control: methods.control });
  const variants = useWatch({ name: "variants", control: methods.control });

  const onSubmit = (value: FormData) => {};

  return (
    <div className="w-screen h-screen flex items-center justify-center">
      <Helmet>
        <title>New Order | TastyFood</title>
      </Helmet>
      <div className="flex flex-col w-full max-w-md px-4 py-8 bg-white rounded-lg shadow sm:px-6 md:px-8 lg:px-10">
        <div className="self-center text-xl font-light text-gray-600 sm:text-2xl">
          What do you like to eat?
        </div>
        <div className="mt-8">
          <FormProvider {...methods}>
            <form noValidate onSubmit={methods.handleSubmit(onSubmit)}>
              <Input
                required
                className="mb-2"
                name="mealName"
                placeholder="Pizza Cipola"
                label="Mealname"
              />
              <div className="flex">
                <Input
                  className="flex-none"
                  name="variant"
                  placeholder="Big"
                  label="Variant"
                />
                <div className="p-6">
                  <Button
                    onClick={() => {
                      if (!variant) {
                        return;
                      }

                      methods.setValue("variants", [...variants, variant]);
                      methods.setValue("variant", "");
                    }}
                    className="flex-auto"
                    type="button"
                  >
                    Add
                  </Button>
                </div>
              </div>
              <div className="flex gap-1 mb-4">
                {variants?.map((variant, index) => (
                  <Badge
                    onClick={() => {
                      const newVariants = variants.filter(function (_, i) {
                        return i !== index;
                      });

                      methods.setValue("variants", newVariants);
                    }}
                  >
                    {variant}
                  </Badge>
                ))}
              </div>
              <Button className="mb-2">Submit and Send</Button>
              <Button variant="secondary">Submit and add another one</Button>
            </form>
          </FormProvider>
        </div>
      </div>
    </div>
  );
};
