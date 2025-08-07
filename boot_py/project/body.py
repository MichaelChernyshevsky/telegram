import os

from dotenv import load_dotenv
from telegram import ReplyKeyboardMarkup, Update
from telegram.ext import Application, CommandHandler, MessageHandler, filters, CallbackContext


load_dotenv()


async def start(update: Update, context: CallbackContext):
    # Создаем клавиатуру с кнопками
    menu_buttons = [
        ["📊 Профиль", "⚙ Настройки"],
        ["ℹ Помощь", "📞 Контакты"]
    ]
    
    reply_markup = ReplyKeyboardMarkup(
        menu_buttons, 
        resize_keyboard=True,  # Кнопки подстраиваются под размер
        one_time_keyboard=False  # Меню остается после нажатия
    )
    
    await update.message.reply_text(
        "🏠 Главное меню:",
        reply_markup=reply_markup
    )

async def handle_menu_buttons(update: Update, context: CallbackContext):
    text = update.message.text
    if text == "📊 Профиль":
        await update.message.reply_text("Здесь будет ваш профиль...")
    elif text == "⚙ Настройки":
        await update.message.reply_text("Раздел настроек...")
    elif text == "ℹ Помощь":
        await update.message.reply_text("Справка по использованию бота...")
    elif text == "📞 Контакты":
        await update.message.reply_text("Наши контакты: @support")




def main():

    token = os.getenv("BOT_TOKEN")
    if not token:
        raise ValueError("Не найден BOT_TOKEN в .env файле")


    """Запуск бота"""
    # Создаем Application вместо Updater
    app = Application.builder().token(token).build()
    
    # Регистрируем обработчики

    
    app.add_handler(CommandHandler("start", start))
    app.add_handler(MessageHandler(filters.TEXT, handle_menu_buttons))

    print("🤖 Бот запущен и ожидает сообщений...")
    app.run_polling()
