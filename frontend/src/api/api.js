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
        try{
            const response = await axios.post(baseUrl + '/login', {
                email: email,
                password: password
            }, {
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded"
                }
            });
            return response;
        } catch (error) {
            console.error('Error fetching login:', error);
            throw error;
        }
    }
}
