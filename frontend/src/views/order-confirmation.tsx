import React from "react";
import { FormProvider, useForm } from "react-hook-form";
import { useQueryClient } from "react-query";
import { useNavigate } from "react-router-dom";
import { Button } from "../components/button/button";
import { OrderList } from "../components/order-list/order-list";
import { Routes } from "../enums/routes.enum";
import { useOrdersByUser } from "../hooks/ordersByUser.hook";
import OrderService from "../services/order.service";

export interface Props {}

/**
 * An OrderConfirmation React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const OrderConfirmation: React.FC<Props> = () => {
  const { data: result } = useOrdersByUser();
  const queryClient = useQueryClient();
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
              onEditClick: () => {
                navigate(Routes.EDIT_ORDER.replace(":name", name ?? ""));
              },
              onDeleteClick: async () => {
                const orderService = new OrderService();

                if (!name) return;

                await orderService.deleteOrder(name);

                await queryClient.invalidateQueries(["orders-by-user"]);
              },
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
