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
     * @returns {Promise<{ ID: string, Nickname: string, Avatar: string, Email: string }>} - Информация о пользователе.
     * */
    static async getUserInfo(token){
        try{
            const response = await axios.get(baseUrl + '/user', {
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                }
            });
            return response.data.UserInfo;
        } catch (error) {
            if (axios.isAxiosError(error)) {
                console.error('Error fetching user info:', error);
                if (error.response && error.response.status === 401 ) {
                    const refreshToken = localStorage.getItem('refreshToken');
                    if (refreshToken) {
                        const newToken = await this.refreshToken(refreshToken);
                        if (newToken) {
                            localStorage.setItem("token", newToken);
                            return this.getUserInfo(newToken);
                        }
                    } else {
                        console.error('Refresh token is missing');
                        throw new Error('Refresh token is missing');
                    }
                }
            }
            throw error;
        }
    }

    /**
     * Выполняет запрос на обновление токена.
     * @param {string} refreshToken - Токен пользователя.
     * */
    static async refreshToken(refreshToken) {
        try {
            const response = await axios.post(baseUrl + '/refresh-token', {
                refreshToken: refreshToken
            });
            return response.data.accessToken;
        } catch (error) {
            console.error('Error fetching refresh token:', error);
            throw error;
        }
    }

    /**
     * Выполняет запрос на создание сессии.
     * @param {Object } sessionData - Данные сессии.
     * @param {string} sessionData.owner_id - Идентификатор владельца сессии.
     * @param {boolean} sessionData.editable - Флаг, указывающий, можно ли редактировать сессию.
     * @param {string} sessionData.title - Заголовок сессии.
     * @param {string} sessionData.language - Язык сессии.
     * @param {number} sessionData.max_users - Максимальное количество пользователей в сессии.
     */
    static async createSession(sessionData){
        try {
            const response = await axios.post(baseUrl + '/session', {
                owner_id: sessionData.owner_id,
                editable: sessionData.editable,
                title: sessionData.title,
                language: sessionData.language,
                max_users: sessionData.max_users
            });
            return response.data;
        } catch (error) {
            console.error('Error fetching create session:', error);
            throw error;
        }
    }


    /**
     *  Выполняет запрос на получение информации о сессии.
     * @param {string} sessionId - Идентификатор сессии.
     */
    static async getSessionInfo(sessionId) {
        return await axios.get(baseUrl + '/session', {
            params: {
                session_id: sessionId,
            }
        })
            .then(response => response.data)
            .catch(error => console.log('Error fetching session info', error));
    }
}
