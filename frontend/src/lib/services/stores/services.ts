import { env } from "$lib";
import type { CookieInfo } from "$lib/helpers/helpers";
import { uploadFiles } from "../products/services";

export async function editStore(
  info: CookieInfo,
  data: FormData,
  update: boolean = false
): Promise<any> {
  var bo: Record<string, any> = {
    image: data.get("image") ?? "",
    name: String(data.get("name")),
    contact_info: {
      email: data.get("email"),
      phone_number: data.get("phone_number"),
      facebook: data.get("facebook"),
      twitter: data.get("twitter"),
      instagram: data.get("instagram"),
    },
  };

  // images
  var files = data.getAll("images");
  if (files.length > 0) {
    const res1 = await uploadFiles(info, data);
    if (res1.status !== 200) {
      return res1;
    }

    const resData1 = await res1.json();
    bo.image = resData1.images[0];
  }

  //body

  if (update) {
    bo.store_id = data.get("store_id");
  }

  const dt = JSON.stringify(bo);
  const res = await fetch(`${env.VITE_URL}/api/v1/stores/create`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${info.token}`,
    },
    body: dt,
  });
  

  return res;
}
