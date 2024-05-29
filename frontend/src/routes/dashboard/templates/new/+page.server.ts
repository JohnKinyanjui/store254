import { getCookieInfo } from '$lib/helpers/helpers';
import { githubLinkCallback, getAvailableRepos, githubLink } from '$lib/services/integrations/github_services';
import { createTemplate } from '$lib/services/templates/services.js';
import { redirect } from '@sveltejs/kit';

export async function load({ url, cookies }) {
    const params = url.searchParams.size;
    if (params !== 0) {
        const res = await githubLinkCallback(getCookieInfo(cookies), {
            installationId: String(url.searchParams.get("installation_id")),
            code: String(url.searchParams.get("code")),
            state: String(url.searchParams.get("state")),
        })
        if (res === "success") {
            throw redirect(301, "/dashboard/templates/new");
        }
    }

    return {
        repositories: await getAvailableRepos(getCookieInfo(cookies))
    }
}


export const actions = {
    create: async function ({ cookies, request }) {
        const data = await request.formData();
        let res = await createTemplate(getCookieInfo(cookies), data, false)
        if (res.status === 200) {
            throw redirect(303, "/dashboard/templates");
        } else {
            return { error: "unable to upload product" };
        }
    },
    githubLink: async ({ cookies }) => {
        console.log("cliclex");

        const url = await githubLink(getCookieInfo(cookies))
        throw redirect(301, url);
    },
}