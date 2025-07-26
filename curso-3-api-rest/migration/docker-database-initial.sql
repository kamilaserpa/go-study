create table personalidades(
        id serial primary key,
        nome varchar,
        historia varchar
    );
    -- Inserindo dados
    INSERT INTO personalidades(nome, historia) VALUES
    ('José de Alencar', 'Romancista, dramaturgo e político, autor de "O Guarani" e "Iracema'),
    ('Fagner', 'Cantor e compositor, conhecido por sua música regional e romântica.'),
    ('Rachel de Queiroz', 'Rachel de Queiroz foi uma escritora pioneira no cenário literário brasileiro. Ela foi a primeira mulher a ser eleita para a Academia Brasileira de Letras.');
