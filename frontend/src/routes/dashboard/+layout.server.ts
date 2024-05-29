import { env } from '$env/dynamic/private';
import { redirect } from '@sveltejs/kit';
import { getCookieInfo, type CookieInfo } from '../../lib/helpers/helpers';
import type { Profile } from '../../models/profile';
import type { LayoutServerLoad } from './$types';


export const load: LayoutServerLoad = async ({ cookies }) => {
    const profile = await _getProfile(getCookieInfo(cookies));
    return {
        profile: profile,
        currentStore: getCookieInfo(cookies).current_store
    };
};

async function _getProfile(cookie: CookieInfo): Promise<Profile> {


    var res = await fetch(`${env.VITE_URL}/api/v1/users/profile`, {
        headers: {
            Authorization: `Bearer ${cookie.token}`
        }
    })



    let profile: Profile = await res.json()
    return profile
}


