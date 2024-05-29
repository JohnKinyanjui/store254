import { env } from "$lib";
import { getDashboardData } from "$lib/services/analytic/analyticService.js";
import type { Store } from "$lib/services/stores/models.js";
import { redirect } from "@sveltejs/kit";
import { getCookieInfo, type CookieInfo } from "$lib/helpers/helpers.js";

export async function load({ cookies }) {
  const cookieInfo = getCookieInfo(cookies);
  const stores = await _getStores(getCookieInfo(cookies));

  if (cookieInfo.current_store.length === 0) {
    throw redirect(303, "/profile");
  } else {
    if (stores.filter((e) => e.id === cookieInfo.current_store).length === 0) {
      throw redirect(303, "/profile");
    }
  }

  return {
    url: env.VITE_URL,
    cookieInfo: cookieInfo,
    dashboard: await getDashboardData(cookieInfo),
    stores: stores,
  };
}

async function _getStores(cookie: CookieInfo) {
  const res = await fetch(`${env.VITE_URL}/api/v1/stores/me?skip=0`, {
    headers: {
      Authorization: `Bearer ${cookie.token}`,
    },
  });

  const stores: Store[] = await res.json();

  return stores;
}
