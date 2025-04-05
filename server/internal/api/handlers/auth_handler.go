package handlers

import (
    "net/http"
    "strings"
    "time"

    "clinica_server/config"
    "clinica_server/internal/api/models"
    "clinica_server/internal/security"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
    DB  *gorm.DB
    Cfg *config.Config
}

func NewAuthHandler(db *gorm.DB, cfg *config.Config) *AuthHandler {
    return &AuthHandler{DB: db, Cfg: cfg}
}

// Login faz a autenticação usando email (ou username) + senha
func (h *AuthHandler) Login(c *gin.Context) {
    var req models.LoginRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Dados de login inválidos",
            "data":    err.Error(),
        })
        return
    }

    var usuario models.Usuario

    // Login usando email (caso queira permitir também username, adapte a consulta)
    result := h.DB.Where("LOWER(email) = ?", strings.ToLower(req.Email)).First(&usuario)
    if result.Error != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "success": false,
            "message": "Falha na autenticação",
            "error":   "Usuário não encontrado",
        })
        return
    }

    // Comparar senha
    if err := bcrypt.CompareHashAndPassword([]byte(usuario.SenhaHash), []byte(req.Senha)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "success": false,
            "message": "Falha na autenticação",
            "error":   "Senha inválida",
        })
        return
    }

    // Atualiza último login
    agora := time.Now()
    h.DB.Model(&usuario).Update("last_login", agora)

    // Gera os tokens
    accessToken, refreshToken, err := security.GenerateTokens(usuario.ID, usuario.Perfil, h.Cfg)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao gerar token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Login efetuado com sucesso",
        "access_token":  accessToken,
        "refresh_token": refreshToken,
        "usuario": usuario.ToResponse(),
    })
}