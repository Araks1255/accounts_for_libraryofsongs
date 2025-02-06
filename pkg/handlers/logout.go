package handlers

import "github.com/gin-gonic/gin"

func (h handler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "Вы успешно вышли из аккаунта"})
}
