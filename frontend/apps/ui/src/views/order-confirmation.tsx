import React, { useState } from "react";
import { Helmet } from "react-helmet";
import { FormProvider, useForm } from "react-hook-form";
import { useQueryClient } from "react-query";
import { Button } from "../components/button/button";
import { CreateOrderView } from "../components/create-order/create-order";
import { OrderList } from "../components/order-list/order-list";
import { Queries } from "../enums/queries.enum";
import { useOrdersByUser } from "../hooks/ordersByUser.hook";
import OrderService from "../services/order.service";
import { useTranslation } from "react-i18next";
import { Timer } from "../components/timer/timer";
import { useIsAdmin } from "../hooks/isAdmin.hook";
import AdminService from "../services/admin.service";
export interface Props {}

/**
 * An OrderConfirmation React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const OrderConfirmation: React.FC<Props> = () => {
  const { data: result } = useOrdersByUser();
  const queryClient = useQueryClient();
  const methods = useForm();
  const { t } = useTranslation("order-confirmation");
  const { data: isAdmin } = useIsAdmin();

  const [open, setOpen] = useState(false);
  const [name, setName] = useState<string | undefined>(undefined);

  return (
    <FormProvider {...methods}>
      <Helmet>
        <title>{t("headline")} | TastyFood</title>
      </Helmet>
      <div className="flex items-center justify-center w-full h-full">
        <Timer />
        <OrderList
          headline={t("headline")}
          items={
            result?.orders?.map(({ name, variants }) => ({
              headline: name ?? "",
              description: variants?.join(", ") ?? "",
              onEditClick: () => {
                setName(name);
                setOpen(true);
              },
              onDeleteClick: async () => {
                const orderService = new OrderService();

                if (!name) return;

                await orderService.deleteOrder(name);

                await queryClient.invalidateQueries(Queries.ORDERS_BY_USER);
              },
              size: 1,
            })) ?? []
          }
        >
          <Button
            className="mt-4"
            type="button"
            onClick={() => {
              setName(undefined);
              setOpen(true);
            }}
          >
            {t("addButton")}
          </Button>
          {isAdmin && (
            <Button
              color="green"
              onClick={async () => {
                const adminService = new AdminService();
                await adminService.next();
              }}
              className="mt-2"
            >
              {t("endPoll")}
            </Button>
          )}
        </OrderList>
      </div>
      <CreateOrderView
        open={open}
        name={name}
        onClose={() => {
          setOpen(false);
        }}
      />
    </FormProvider>
  );
};
