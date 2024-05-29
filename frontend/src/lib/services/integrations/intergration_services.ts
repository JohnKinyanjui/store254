import { env } from "$lib";
import type { CookieInfo } from "$lib/helpers/helpers";

export async function getOauthIntergrations(
  cookie: CookieInfo
): Promise<AuthenticationConfig[]> {
  const res = await fetch(
    `${env.VITE_URL}/api/v1/cloud/oauth/all?id=${cookie.current_store}`,
    {
      method: "GET",
      headers: {
        Authorization: `Bearer ${cookie.token}`,
      },
    }
  );

  const r = await res.json();

  const oauths: AuthenticationConfig[] = r;
  return oauths;
}

export async function createIntergrateOauth(
  cookie: CookieInfo,
  data: FormData
): Promise<any> {
  let creds: Record<string, string> = {};
  const authType = data.get("auth_type");

  switch (authType) {
    case "google":
      creds = {
        consumer_key: String(data.get("consumer_key") ?? ""),
        consumer_secret: String(data.get("consumer_secret") ?? ""),
        redirect_url: String(data.get("redirect_url")),
      };
      break;

    default:
      break;
  }
  let body: Record<string, any> = {
    store_id: cookie.current_store,
    auth_type: String(authType),
    credentials: creds,
  };

  const res = await fetch(
    `${env.VITE_URL}/api/v1/cloud/oauth/create?id=${cookie.current_store}`,
    {
      method: "POST",
      headers: {
        Authorization: `Bearer ${cookie.token}`,
      },
      body: JSON.stringify(body),
    }
  );

  const r = await res.json();
  console.log(r);

  return r;
}
