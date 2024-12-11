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
     * Выполняет вход пользователя.
     * @param {string} email - Электронная почта пользователя.
     * @param {string} password - Пароль пользователя.
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

    /**
     * Выполняет регистрацию нового пользователя.
     * @param {string} email - Электронная почта пользователя.
     * @param {string} password - Пароль пользователя.
     */
    static async register(email, password){
        try{
            const response = await axios.post(baseUrl + '/register', {
                email: email,
                password: password
            }, {
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded"
                }
            });
            return response;
        } catch (error) {
            console.error('Error fetching register:', error);
            throw error;
        }
    }

    /**
     * Выполняет запрос на получение информации о пользователе.
     * @param {string} token - Токен пользователя.
     * */
    static async getUserInfo(token){
        return await axios.get(baseUrl + '/user', {
            headers: {
                "Authorization": "Bearer " + token
            }
        })
            .then(response => response.data)
            .catch(error => console.log('Error fetching user info', error));
    }
}
