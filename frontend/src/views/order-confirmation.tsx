import React from "react";
import { FormProvider, useForm } from "react-hook-form";
import { useLocation, useNavigate, useParams } from "react-router-dom";
import { Button } from "../components/button/button";
import { OrderList } from "../components/order-list/order-list";
import { Routes } from "../enums/routes.enum";
import { useOrdersByUser } from "../hooks/ordersByUser.hook";

export interface Props {}

/**
 * An OrderConfirmation React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const OrderConfirmation: React.FC<Props> = () => {
  const { user } = useParams<"user">();
  const { data: result } = useOrdersByUser(user ?? "");
  const navigate = useNavigate();
  const methods = useForm();

  return (
    <FormProvider {...methods}>
      <div className="flex items-center justify-center w-full h-full">
        <OrderList
          headline="My Orders"
          items={
            result?.orders?.map(({ name, variants }) => ({
              headline: name ?? "",
              description: variants?.join(", ") ?? "",
              size: 1,
            })) ?? []
          }
        >
          <Button
            className="mt-4"
            type="button"
            onClick={() => {
              navigate(Routes.NEW_ORDER);
            }}
          >
            Add
          </Button>
        </OrderList>
      </div>
    </FormProvider>
  );
};
