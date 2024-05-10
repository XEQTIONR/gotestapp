import axios from 'axios'

export async function useData(to, router) {
    let data = {}
    if (window.originURL == to && window.valid) {
        // dont make api call
        data = window.apiData
        window.valid = false
    } else {
        try {
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