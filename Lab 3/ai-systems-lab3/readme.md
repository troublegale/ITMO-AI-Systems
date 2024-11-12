# Lab 3. Linear regression

Было построено 3 модели с различными наборами признаков.
Первая включает все признаки из предоставленного датасета,
вторая исключает признак `extracurricular_activities`,
третья добавляет к признакам из датасета дополнительный признак
`study_sleep_ratio`, показывающий отношение признаков
`hours_studied` и `sleep_hours`.

В ходе сравнения моделей третья модель была определена как
самая производительная, опираясь на значение
коэффициента детерминации, который оказался самым
высоким именно у неё.

Eng: there are 3 models. 1st one is normal, 2nd one has
`extracurricular_activities` removed, 3rd one has 
`study_sleep_ratio` added. As a result, the 3rd model
turned out to be the best one based on the
determination coefficient.