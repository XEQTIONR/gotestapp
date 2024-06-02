import axios from 'axios'

export async function useData(to, router) {
    let data = {}
    if (window.originURL == to) {
        // dont make api call
        data = window.apiData
    } else {
        try {
            axios.defaults.withXSRFToken = true;
            const response = await axios.get(to)
            data = response.data

            if (data.errors) {
                data.errors = JSON.parse(data.errors)
            }
        } catch (e) {
            if (e.response.status === 401) {
                router.replace({ path: '/login'})
            }
        }
    }

    return data 
}