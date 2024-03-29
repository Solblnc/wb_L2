<h1>Фабричный метод</h1>
<p>
    Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в суперклассе,
    позволяя подклассам изменять тип создаваемых объектов.
</p>
<h2>Применимость</h2>
<ul>
    <li>
        Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код
    </li>
    <li>
        <p> Когда вы хотите дать возможность пользователям расширять части вашего фреймворка или библиотеки</p>
    </li>
    <li>
        <p>
            Когда вы хотите экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых.
            Такая проблема обычно возникает при работе с тяжёлыми ресурсоёмкими объектами, такими, как подключение к базе данных, файловой системе и т. д.
        </p>
    </li>
</ul>
<h2>Плюсы и минусы</h2>
<ul>
    <b>Плюсы</b>
    <li>Избавляет класс от привязки к конкретным классам продуктов</li>
    <li>Выделяет код производства продуктов в одно место, упрощая поддержку кода</li>
    <li>Упрощает добавление новых продуктов в программу</li>
    <li>Реализует принцип открытости/закрытости</li>
    </br>
    <b>Минусы</b>
    <li>Может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать свой подкласс создателя</li>
</ul>
<h2>Реальные примеры</h2>
<p>
    В Go невозможно реализовать классический вариант паттерна Фабричный метод, поскольу в языке отсутствуют возможности ООП, в том числе классы и наследственность.
    Несмотря на это, мы все же можем реализовать базовую версию этого паттерна — Простая фабрика.
</p>
<p>В этом примере мы будем создавать разные типы оружия при помощи структуры фабрики.</p>
<p>
    Сперва, мы создадим интерфейс <code>iGun</code>, который определяет все методы будущих пушек.
    Также имеем структуру <code>gun</code> (пушка), которая применяет интерфейс <code>iGun</code>.
    Две конкретных пушки — <code>ak47</code> и <code>musket</code> — обе включают в себя структуру <code>gun</code> и не напрямую реализуют все методы от <code>iGun</code>.
</p>
<p>
    <code>gunFactory</code> служит фабрикой, которая создает пушку нужного типа в зависимости от аргумента на входе.
    Вместо прямого взаимодействия с объектами <code>ak47</code> или <code>musket</code>, он создает экземпляры различного оружия при помощи <code>gunFactory</code>,
    используя для контроля изготовления только параметры в виде строк.
</p>