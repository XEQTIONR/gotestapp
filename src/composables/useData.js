export async function useData(keys = []) {
    
    let data = {}
    // this.busy = true
    if (window.originURL == window.location.href && window.valid) {
        // dont make api call
        data = window.apiData
        window.valid = false
    } else {

        const response = await axios.get(window.location.href)
        data = response.data
    }

    return data 
}