import {
  createPaymentIntegration,
  createPaymentMethod,
  deletePaymentMethod,
  getAvailablePaymentMethods,
  getPaymentMethods,
} from "$lib/services/payments/services";
import { getCookieInfo } from "$lib/helpers/helpers";

export async function load({ url, cookies }) {
  const search = url.searchParams.get("search");
  return {
    methods: await getPaymentMethods(
      getCookieInfo(cookies),
      search == null ? "" : search
    ),
    available: await getAvailablePaymentMethods(getCookieInfo(cookies)),
  };
}

export let actions = {
  create: async function ({ request, cookies }) {
    const data = await request.formData();

    let msg = await createPaymentMethod(getCookieInfo(cookies), data, false);
    if (msg == "success") {
      return { success: "success" };
    }

    return { error: msg };
  },

  update: async function ({ request, cookies }) {
    const data = await request.formData();

    let msg = await createPaymentMethod(getCookieInfo(cookies), data, true);
    if (msg == "success") {
      return { success: "success" };
    }

    return { error: msg };
  },

  intergrate: async function ({ request, cookies }) {
    const data = await request.formData();

    let msg = await createPaymentIntegration(getCookieInfo(cookies), data);
    if (msg == "success") {
      return { success: "success" };
    }

    return { error: msg };
  },
  delete: async function ({ request, cookies }) {
    const data = await request.formData();

    let msg = await deletePaymentMethod(getCookieInfo(cookies), data);
    if (msg == "success") {
      return { success: "success" };
    }

    return { error: msg };
  },
};
