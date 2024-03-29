<h1>Состояние</h1>
<p>
    Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего состояния.
    Извне создаётся впечатление, что изменился класс объекта.
</p>
<p>Паттерн Состояние предлагает создать отдельные классы для каждого состояния, в котором может пребывать объект, а затем вынести туда поведения, соответствующие этим состояниям.</p>
<p>
    Вместо того, чтобы хранить код всех состояний, первоначальный объект, называемый контекстом, будет содержать ссылку на один из объектов-состояний и делегировать ему работу, зависящую от состояния.
</p>
<h2>Применимость</h2>
<ul>
    <li>
        Когда у есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния,
        причём типов состояний много, и их код часто меняется
    </li>
    <li>
        <p>
            Когда код класса содержит множество больших, похожих друг на друга, условных операторов,
            которые выбирают поведения в зависимости от текущих значений полей класса.
        </p>
    </li>
    <li>
        Когда сознательно используется табличная машина состояний, построенную на условных операторах,
        но вынуждены мириться с дублированием кода для похожих состояний и переходов.
    </li>
</ul>
<h2>Плюсы и минусы</h2>
<ul>
    <b>Плюсы</b>
    <li>Избавляет от множества больших условных операторов машины состояний</li>
    <li>Концентрирует в одном месте код, связанный с определённым состоянием.</li>
    <li>Упрощает код контекста</li>
    </br>
    <b>Минусы</b>
    <li>Может неоправданно усложнить код, если состояния мало и они редко меняются</li>
</ul>
<h2>Реальные примеры</h2>
<p>
    Давайте применим паттерн проектирования Состояние в контексте торговых автоматов.
    Для упрощения задачи представим, что торговый автомат может выдавать только один товар.
    Также представим, что автомат может пребывать только в одном из четырех состояний:
</p>
<ul>
    <li>hasItem (имеетПредмет)</li>
    <li>noItem (неИмеетПредмет)</li>
    <li>itemRequested (выдаётПредмет)</li>
    <li>hasMoney (получилДеньги)</li>
</ul>
<p>
    Торговый автомат может иметь различные действия. Опять-таки, для простоты оставим только четыре из них:
</p>
<ul>
    <li>Выбрать предмет</li>
    <li>Добавить предмет</li>
    <li>Ввести деньги</li>
    <li>Выдать предмет</li>
</ul>
<p>
    Паттерн Состояние нужно использовать в случаях, когда объект может иметь много различных состояний,
    которые он должен менять в зависимости от конкретного поступившего запроса.
</p>
<p>
    В нашем примере, автомат может быть в одном из множества состояний, которые непрерывно меняются.
    Припустим, что торговый автомат находится в режиме <code>itemRequested</code>. Как только произойдет действие «Ввести деньги»,
    он сразу же перейдет в состояние <code>hasMoney</code>.
</p>
<p>
    В зависимости от состояния торгового автомата, в котором он находится на данный момент, он может по-разному отвечать на одни и те же запросы.
    Например, если пользователь хочет купить предмет, машина выполнит действие, если она находится в режиме <code>hasItemState</code>, и отклонит запрос в режиме <code>noItemState</code>.
</p>
<p>
    Паттерн Состояние нужно использовать в случаях, когда объект может иметь много различных состояний,
    которые он должен менять в зависимости от конкретного поступившего запроса.
</p>