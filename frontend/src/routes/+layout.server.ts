import { env } from '$lib';
import type { Store } from '$lib/services/stores/models';
import { getCookieInfo, type CookieInfo } from '../lib/helpers/helpers';
import type { Profile } from '../models/profile';

export const load = async ({ cookies, url }) => {
  

    const profile = await _getProfile(getCookieInfo(cookies))
    const stores = await _getStores(getCookieInfo(cookies))
    return {
        profile: profile,
        stores: stores,
        err: undefined
    };
};

async function _getProfile(cookie: CookieInfo) {
    if (cookie.token.length === 0) {
        return undefined
    }
    
    var res = await fetch(`${env.VITE_URL}/api/v1/users/profile`, {
        headers: {
            Authorization: `Bearer ${cookie.token}`,
            accept: "application/json"
        }
    })

    if (res.status === 200) {
        let profile: Profile = await res.json()
        return profile
    }

    return undefined
}


async function _getStores(cookie: CookieInfo) {
    const res = await fetch(`${env.VITE_URL}/api/v1/stores/me?skip=0`, {
        headers: {
            Authorization: `Bearer ${cookie.token}`,
            accept: "application/json"

        }
    })

    if (res.status === 200) {
        const stores: Store[] = await res.json();
        return stores
    }


    return []
}