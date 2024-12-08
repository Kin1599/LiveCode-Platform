import axios from "axios"
import baseUrl from "../configs/config";

export default class SendServer{
    //* Здесь можно писать функции, которые взаимодействуют с сервером

    //* Функция для проверки на CORS
    static async getPing(){
        return await axios.get(baseUrl + '/ping')
            .then(response => response.data)
            .catch(error => console.log('Error fetching products', error));
    }

    /**
     * @param {string} email
     * @param {any} password
     */
    static async login(email, password){
        return await axios.post(baseUrl + '/login', {}, {
            params: {
                email: email,
                password: password
            }
        })
            .then(response => response.data)
            .catch(error => console.log('Error fetching login', error));
    }
}
