<h1>Строитель</h1>
Строитель (builder) - порождающий паттерн проектирования, который позволяет поэтапно создавать сложные объекты.
Строитель дает возможность использовать один и тот же код строительства для получения разных представлений объекта.
<h2>Применимость</h2>
<ul>
    <li>Чтобы избавиться от «телескопического конструктора» c большим количеством опциональных параметров</li>
    <li>Когда код должен создавать разные представления какого-то объекта. Например, деревянные и железобетонные дома</li>
    <li>Когда нужно собирать сложные составные объекты, например, деревья Компоновщика</li>
</ul>
<h2>Плюсы и минусы</h2>
<ul>
    <b>Плюсы</b>
    <li>Позволяет создавать продукты пошагово, при этом не позволяет посторонним объектам иметь доступ к конструируему объекту, пока тот не будет полностью готов</li>
    <li>Позволяет использовать один и тот же код для создания различных продуктов</li>
    <li>Изолирует сложный код сборки продукта от его основной бизнес-логики</li>
    </br>
    <b>Минусы</b>
    <li>Усложняет код программы из-за введения дополнительных классов</li>
    <li>Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата</li>
</ul>
<h2>Реальные примеры</h2>
Паттерн Строитель также используется, когда нужный продукт сложный и требует нескольких шагов для построения.
В таких случаях несколько конструкторных методов подойдут лучше, чем один громадный конструктор.
При использовании пошагового построения объектов потенциальной проблемой является выдача клиенту частично построенного нестабильного продукта.
Паттерн "Строитель" скрывает объект до тех пор, пока он не построен до конца.
</br>
</br>
Нижеприведенный код использует разные типы автомобилей (<code>audi</code> и <code>ferrari</code>), которые конструируются с помощью строителей <code>ferrari-builder</code> и <code>audi-builder</code>.
При создании каждого авто используются одинаковые шаги. Для помощи в организации процесса можно использовать директор, хоть это и необязательно.