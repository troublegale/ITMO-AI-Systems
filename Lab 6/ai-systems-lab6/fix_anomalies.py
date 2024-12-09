import pandas as pd
import numpy as np

# Загрузка данных
file_path = 'resources/diabetes.csv'
data = pd.read_csv(file_path)

# Диапазоны для каждой колонки
ranges = {
    'Glucose': (50, 250),
    'BloodPressure': (40, 140),
    'SkinThickness': (10, 80),
    'Insulin': (10, 600),
    'BMI': (10, 70)
}

# Функция для обработки аномальных значений
def handle_outliers_with_median(column, value_range):
    lower, upper = value_range
    # Вычисляем медиану только на данных в пределах допустимого диапазона
    median_value = data[(data[column] >= lower) & (data[column] <= upper)][column].median()
    # Заменяем аномалии на вычисленную медиану
    data[column] = data[column].apply(lambda x: median_value if x < lower or x > upper else x)


def fix_dataset_and_save():
    # Применение функции для указанных колонок
    for column, value_range in ranges.items():
        handle_outliers_with_median(column, value_range)
    # Сохранение очищенного датасета
    cleaned_file_path = 'resources/fixed_diabetes.csv'
    data.to_csv(cleaned_file_path, index=False)

    print(f"Данные очищены от аномалий и сохранены в {cleaned_file_path}")
