SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;


-- --------------------------------------------------------


--
-- テーブルのデータのダンプ `users`
--

INSERT INTO `users` (`id`, `name`, `mail`, `image`, `created_at`) VALUES
(1, 'yuto sekiguchi', 'sekiguchi@gmail.com', 'https://lh3.googleusercontent.com/a-/AOh14GgpQd4iz3eUTS7-vqcFmLv7NRVA954p7bi2S8VFQA=s96-c', '2021-10-26 11:48:30'),
(2, 'nisemura usoshi', 'uso800@nise.lie', 'https://dummyimage.com/400x400/000/fff', '2021-10-26 21:57:09');


-- --------------------------------------------------------

--
-- テーブルのデータのダンプ `comics`
--

INSERT INTO `comics` (`id`, `name`, `created_at`) VALUES
(1, '進撃の巨人', '2021-09-15 17:26:57'),
(2, '呪術廻戦', '2021-09-15 17:26:57'),
(3, '鬼滅の刃', '2021-09-16 10:30:19'),
(4, 'チェンソーマン', '2021-09-16 10:30:19'),
(5, '僕のヒーローアカデミア', '2021-09-16 10:31:09'),
(6, '約束のネバーランド', '2021-09-16 10:31:09'),
(7, '名探偵コナン', '2021-09-16 10:31:09');


-- --------------------------------------------------------

--
-- テーブルのデータのダンプ `comic_volumes`
--

INSERT INTO `comic_volumes` (`id`, `volume`, `cid`, `all_page`, `ebook_url`, `comic_vol_image`, `created_at`) VALUES
(1, 1, 1, 192, 'https://ebookjapan.yahoo.co.jp/books/121579/A000053595/', 'https://cache2-ebookjapan.akamaized.net/contents/thumb/s/KD596860.jpg?158046119100', '2021-09-15 17:26:57'),
(2, 2, 1, 191, 'https://ebookjapan.yahoo.co.jp/books/121579/A000059601/', 'https://cache2-ebookjapan.akamaized.net/contents/thumb/s/KD658960.jpg?1580461191000', '2021-09-15 17:26:57'),
(3, 1, 2, 196, 'https://ebookjapan.yahoo.co.jp/books/464912/A001899692/', 'https://cache2-ebookjapan.akamaized.net/contents/thumb/s/Q5100038208961.jpg?1622198079000', '2021-09-16 10:30:19'),
(4, 2, 3, 197, 'https://ebookjapan.yahoo.co.jp/books/367324/A001667922/', 'https://cache2-ebookjapan.akamaized.net/contents/thumb/s/F7100016327061.jpg?1527743388000', '2021-09-16 10:30:19'),
(5, 1, 4, 199, 'https://ebookjapan.yahoo.co.jp/books/527520/A002124290/', 'https://cache2-ebookjapan.akamaized.net/contents/thumb/s/U9100054244661.jpg?1555982140000', '2021-09-16 10:31:09'),
(6, 2, 4, 199, 'https://ebookjapan.yahoo.co.jp/books/527520/A002168796/', 'https://cache2-ebookjapan.akamaized.net/contents/thumb/s/L8100058530961.jpg?1556275776000', '2021-09-16 10:31:09'),
(7, 1, 6, 199, 'https://ebookjapan.yahoo.co.jp/books/390133/A002731124/', 'https://cache2-ebookjapan.akamaized.net/contents/thumb/s/C4100112192461.jpg?1634548401000', '2021-09-16 10:30:19'),
(8, 1, 5, 196, 'https://ebookjapan.yahoo.co.jp/books/264558/A000338425/', 'https://cache2-ebookjapan.akamaized.net/contents/thumb/s/RR890560.jpg?1574330432000', '2021-09-16 10:30:19'),
(9, 2, 6, 199, 'https://ebookjapan.yahoo.co.jp/books/390133/A002731125/', 'https://cache2-ebookjapan.akamaized.net/contents/thumb/s/Y8100020837461.jpg?1573214139000', '2021-09-16 10:31:09'),
(10, 1, 7, 183, 'https://ebookjapan.yahoo.co.jp/books/119961/A000049610/', 'https://cache2-ebookjapan.akamaized.net/contents/thumb/s/SX214360.jpg?1576568623000', '2021-09-16 10:31:09'),
(11, 1, 3, 198, 'https://ebookjapan.yahoo.co.jp/books/367324/A001650413/', 'https://cache2-ebookjapan.akamaized.net/contents/thumb/s/O4100014595261.jpg?1527743388000', '2021-09-16 10:32:19')
;


-- --------------------------------------------------------

--
-- テーブルのデータのダンプ characters`
--

INSERT INTO `characters` (`id`, `name`, `created_at`) VALUES
(1, 'エマ', '2021-10-27 07:19:51'),
(2, 'ノーマン', '2021-10-27 07:19:51'),
(3, 'コニー', '2021-10-27 07:20:05'),
(4, '虎杖悠仁', '2021-10-27 07:20:05'),
(5, '伏黒恵', '2021-10-27 07:20:05'),
(6, '宿儺', '2021-10-28 01:20:10'),
(7, '釘崎野薔薇', '2021-10-28 01:20:10'),
(8, '工藤新一', '2021-10-28 02:24:21');


-- --------------------------------------------------------

--
-- テーブルのデータのダンプ `scenes`
--

INSERT INTO `scenes` (`id`, `uid`, `cvid`, `page`, `scene`, `emotion`, `detail_scene`,`comment`,`created_at`) VALUES
(1, 1, 7, 42, '死亡シーン', '悲しい', 'コニーの死亡シーン', '大好きなコニーが…', '2021-10-27 07:19:51'),
(2, 1, 3, 54, '戦闘シーン', '怖い', '宿儺初登場シーン', '宿儺かっこいい', '2021-10-27 07:19:51'),
(3, 1, 3, 112, 'ギャグシーン', '面白い', '釘崎野薔薇初登場シーン', '', '2021-10-27 07:20:05'),
(4, 2, 10, 34, '衝撃シーン', '怖い', '目が覚めたら体が縮んでしまっていたシーン', '新一ーーーーーーーーーーーーーーーー', '2021-10-27 07:20:05'),
(5, 1, 3, 166, '戦闘シーン', '怖い', '初の特級呪霊との遭遇シーン', '', '2021-10-27 07:20:07'),
(6, 2, 3, 54, '衝撃シーン', '怖い', '虎杖が宿儺になるシーン', '怖いけどかっこいい！', '2021-11-16 07:19:51');


-- --------------------------------------------------------

--
-- テーブルのデータのダンプ `appearance_characters`
--

INSERT INTO `appearance_characters` (`id`, `scid`, `cvid`, `chid`, `created_at`) VALUES
(1, 1, 7, 1, '2021-10-27 07:19:51'),
(2, 1, 7, 2, '2021-10-27 07:19:51'),
(3, 1, 7, 3, '2021-10-27 07:20:05'),
(4, 2, 3, 4, '2021-10-27 07:20:05'),
(5, 2, 3, 5, '2021-10-27 07:32:05'),
(6, 2, 3, 6, '2021-10-27 07:32:05'),
(7, 3, 3, 7, '2021-10-27 07:32:05'),
(8, 3, 3, 4, '2021-10-27 07:32:05'),
(9, 3, 3, 5, '2021-10-27 07:32:05'),
(10, 4, 10, 8, '2021-10-27 07:32:05'),
(11, 5, 3, 5, '2021-10-27 07:32:05'),
(12, 5, 3, 4, '2021-10-27 07:32:05'),
(13, 2, 3, 4, '2021-10-28 07:20:05'),
(14, 2, 3, 5, '2021-10-28 07:32:05'),
(15, 2, 3, 6, '2021-10-28 07:32:05');

-- --------------------------------------------------------

--
-- テーブルのデータのダンプ `readings`
--

INSERT INTO `readings` (`id`, `uid`, `cid`, `created_at`) VALUES
(1, 1, 1, '2021-10-27 07:19:51'),
(2, 1, 2, '2021-10-27 07:19:51'),
(3, 2, 2, '2021-10-27 07:20:05'),
(4, 2, 4, '2021-10-27 07:20:05');