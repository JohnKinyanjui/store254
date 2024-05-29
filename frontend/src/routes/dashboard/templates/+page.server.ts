import { getCookieInfo } from '$lib/helpers/helpers.js'
import { getAvailableRepos, githubLink, githubLinkCallback } from '$lib/services/integrations/github_services';
import { getTemplates } from '$lib/services/templates/services.js';

export async function load({ cookies, url }) {
    return {
        templates: await getTemplates(getCookieInfo(cookies))
    }
}


export let actions = {
}