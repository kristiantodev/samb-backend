-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 07 Des 2024 pada 17.41
-- Versi server: 5.7.21-log
-- Versi PHP: 8.0.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `samb`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `mastercustomer`
--

CREATE TABLE `mastercustomer` (
  `CustomerPK` int(11) NOT NULL,
  `CustomerName` varchar(255) CHARACTER SET utf8 NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `mastercustomer`
--

INSERT INTO `mastercustomer` (`CustomerPK`, `CustomerName`) VALUES
(1, 'Customer X'),
(2, 'Customer Y'),
(3, 'Customer Z'),
(4, 'Customer W'),
(5, 'Customer V');

-- --------------------------------------------------------

--
-- Struktur dari tabel `masterproduct`
--

CREATE TABLE `masterproduct` (
  `ProductPK` int(11) NOT NULL,
  `ProductName` varchar(255) CHARACTER SET utf8 NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `masterproduct`
--

INSERT INTO `masterproduct` (`ProductPK`, `ProductName`) VALUES
(1, 'Product 1'),
(2, 'Product 2'),
(3, 'Product 3'),
(4, 'Product 4'),
(5, 'Product 5');

-- --------------------------------------------------------

--
-- Struktur dari tabel `mastersupplier`
--

CREATE TABLE `mastersupplier` (
  `SupplierPK` int(11) NOT NULL,
  `SupplierName` varchar(255) CHARACTER SET utf8 NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `mastersupplier`
--

INSERT INTO `mastersupplier` (`SupplierPK`, `SupplierName`) VALUES
(1, 'Supplier A'),
(2, 'Supplier B'),
(3, 'Supplier C'),
(4, 'Supplier D'),
(5, 'Supplier E');

-- --------------------------------------------------------

--
-- Struktur dari tabel `masterwarehouse`
--

CREATE TABLE `masterwarehouse` (
  `WhsPK` int(11) NOT NULL,
  `WhsName` varchar(255) CHARACTER SET utf8 NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `masterwarehouse`
--

INSERT INTO `masterwarehouse` (`WhsPK`, `WhsName`) VALUES
(1, 'Warehouse A'),
(2, 'Warehouse B'),
(3, 'Warehouse C'),
(4, 'Warehouse D'),
(5, 'Warehouse E');

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaksipenerimaanbarangdetail`
--

CREATE TABLE `transaksipenerimaanbarangdetail` (
  `TrxInDPK` int(11) NOT NULL,
  `TrxInIDF` int(11) NOT NULL,
  `TrxInDProductIdf` int(11) NOT NULL,
  `TrxInDQtyDus` int(11) NOT NULL,
  `TrxInDQtyPcs` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `transaksipenerimaanbarangdetail`
--

INSERT INTO `transaksipenerimaanbarangdetail` (`TrxInDPK`, `TrxInIDF`, `TrxInDProductIdf`, `TrxInDQtyDus`, `TrxInDQtyPcs`) VALUES
(1, 1, 1, 10, 100),
(2, 1, 2, 5, 50),
(3, 2, 3, 8, 80),
(4, 2, 4, 12, 120),
(5, 3, 5, 15, 150),
(6, 6, 1, 10, 100),
(7, 6, 2, 5, 50);

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaksipenerimaanbarangheader`
--

CREATE TABLE `transaksipenerimaanbarangheader` (
  `TrxInPK` int(11) NOT NULL,
  `TrxInNo` varchar(255) CHARACTER SET utf8 NOT NULL,
  `WhsIdf` int(11) NOT NULL,
  `TrxInDate` varchar(100) NOT NULL,
  `TrxInSuppIdf` int(11) NOT NULL,
  `TrxInNotes` varchar(255) CHARACTER SET utf8 DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `transaksipenerimaanbarangheader`
--

INSERT INTO `transaksipenerimaanbarangheader` (`TrxInPK`, `TrxInNo`, `WhsIdf`, `TrxInDate`, `TrxInSuppIdf`, `TrxInNotes`) VALUES
(1, 'TRXIN001', 1, '2024-12-01', 1, 'Received items for warehouse A'),
(2, 'TRXIN002', 2, '2024-12-02', 2, 'Received items for warehouse B'),
(3, 'TRXIN003', 3, '2024-12-03', 3, 'Received items for warehouse C'),
(4, 'TRXIN004', 4, '2024-12-04', 4, 'Received items for warehouse D'),
(5, 'TRXIN005', 5, '2024-12-05', 5, 'Received items for warehouse E'),
(6, 'TRXIN001', 1, '2024-12-01', 1, 'Received items for warehouse A');

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaksipengeluaranbarangdetail`
--

CREATE TABLE `transaksipengeluaranbarangdetail` (
  `TrxOutDPK` int(11) NOT NULL,
  `TrxOutIDF` int(11) NOT NULL,
  `TrxOutDProductIdf` int(11) NOT NULL,
  `TrxOutDQtyDus` int(11) NOT NULL,
  `TrxOutDQtyPcs` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `transaksipengeluaranbarangdetail`
--

INSERT INTO `transaksipengeluaranbarangdetail` (`TrxOutDPK`, `TrxOutIDF`, `TrxOutDProductIdf`, `TrxOutDQtyDus`, `TrxOutDQtyPcs`) VALUES
(1, 1, 1, 5, 50),
(2, 1, 2, 3, 30),
(3, 2, 3, 7, 70),
(4, 2, 4, 10, 100),
(5, 3, 5, 12, 120),
(8, 7, 1, 5, 50),
(9, 7, 2, 3, 30);

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaksipengeluaranbarangheader`
--

CREATE TABLE `transaksipengeluaranbarangheader` (
  `TrxOutPK` int(11) NOT NULL,
  `TrxOutNo` varchar(255) CHARACTER SET utf8 NOT NULL,
  `WhsIdf` int(11) NOT NULL,
  `TrxOutDate` date NOT NULL,
  `TrxOutSuppIdf` int(11) NOT NULL,
  `TrxOutNotes` varchar(255) CHARACTER SET utf8 DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `transaksipengeluaranbarangheader`
--

INSERT INTO `transaksipengeluaranbarangheader` (`TrxOutPK`, `TrxOutNo`, `WhsIdf`, `TrxOutDate`, `TrxOutSuppIdf`, `TrxOutNotes`) VALUES
(1, 'TRXOUT001', 1, '2024-12-01', 1, 'Sent items from warehouse A'),
(2, 'TRXOUT002', 2, '2024-12-02', 2, 'Sent items from warehouse B'),
(3, 'TRXOUT003', 3, '2024-12-03', 3, 'Sent items from warehouse C'),
(4, 'TRXOUT004', 4, '2024-12-04', 4, 'Sent items from warehouse D'),
(5, 'TRXOUT005', 5, '2024-12-05', 5, 'Sent items from warehouse E'),
(6, 'TRXOUT001', 1, '2024-12-01', 1, 'Sent items from warehouse A'),
(7, 'TRXOUT001', 1, '2024-12-01', 1, 'Sent items from warehouse A');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `mastercustomer`
--
ALTER TABLE `mastercustomer`
  ADD PRIMARY KEY (`CustomerPK`);

--
-- Indeks untuk tabel `masterproduct`
--
ALTER TABLE `masterproduct`
  ADD PRIMARY KEY (`ProductPK`);

--
-- Indeks untuk tabel `mastersupplier`
--
ALTER TABLE `mastersupplier`
  ADD PRIMARY KEY (`SupplierPK`);

--
-- Indeks untuk tabel `masterwarehouse`
--
ALTER TABLE `masterwarehouse`
  ADD PRIMARY KEY (`WhsPK`);

--
-- Indeks untuk tabel `transaksipenerimaanbarangdetail`
--
ALTER TABLE `transaksipenerimaanbarangdetail`
  ADD PRIMARY KEY (`TrxInDPK`),
  ADD KEY `TrxInIDF` (`TrxInIDF`),
  ADD KEY `TrxInDProductIdf` (`TrxInDProductIdf`);

--
-- Indeks untuk tabel `transaksipenerimaanbarangheader`
--
ALTER TABLE `transaksipenerimaanbarangheader`
  ADD PRIMARY KEY (`TrxInPK`),
  ADD KEY `WhsIdf` (`WhsIdf`),
  ADD KEY `TrxInSuppIdf` (`TrxInSuppIdf`);

--
-- Indeks untuk tabel `transaksipengeluaranbarangdetail`
--
ALTER TABLE `transaksipengeluaranbarangdetail`
  ADD PRIMARY KEY (`TrxOutDPK`),
  ADD KEY `TrxOutIDF` (`TrxOutIDF`),
  ADD KEY `TrxOutDProductIdf` (`TrxOutDProductIdf`);

--
-- Indeks untuk tabel `transaksipengeluaranbarangheader`
--
ALTER TABLE `transaksipengeluaranbarangheader`
  ADD PRIMARY KEY (`TrxOutPK`),
  ADD KEY `WhsIdf` (`WhsIdf`),
  ADD KEY `TrxOutSuppIdf` (`TrxOutSuppIdf`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `mastercustomer`
--
ALTER TABLE `mastercustomer`
  MODIFY `CustomerPK` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT untuk tabel `masterproduct`
--
ALTER TABLE `masterproduct`
  MODIFY `ProductPK` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT untuk tabel `mastersupplier`
--
ALTER TABLE `mastersupplier`
  MODIFY `SupplierPK` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT untuk tabel `masterwarehouse`
--
ALTER TABLE `masterwarehouse`
  MODIFY `WhsPK` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT untuk tabel `transaksipenerimaanbarangdetail`
--
ALTER TABLE `transaksipenerimaanbarangdetail`
  MODIFY `TrxInDPK` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT untuk tabel `transaksipenerimaanbarangheader`
--
ALTER TABLE `transaksipenerimaanbarangheader`
  MODIFY `TrxInPK` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT untuk tabel `transaksipengeluaranbarangdetail`
--
ALTER TABLE `transaksipengeluaranbarangdetail`
  MODIFY `TrxOutDPK` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT untuk tabel `transaksipengeluaranbarangheader`
--
ALTER TABLE `transaksipengeluaranbarangheader`
  MODIFY `TrxOutPK` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `transaksipenerimaanbarangdetail`
--
ALTER TABLE `transaksipenerimaanbarangdetail`
  ADD CONSTRAINT `transaksipenerimaanbarangdetail_ibfk_1` FOREIGN KEY (`TrxInIDF`) REFERENCES `transaksipenerimaanbarangheader` (`TrxInPK`),
  ADD CONSTRAINT `transaksipenerimaanbarangdetail_ibfk_2` FOREIGN KEY (`TrxInDProductIdf`) REFERENCES `masterproduct` (`ProductPK`);

--
-- Ketidakleluasaan untuk tabel `transaksipenerimaanbarangheader`
--
ALTER TABLE `transaksipenerimaanbarangheader`
  ADD CONSTRAINT `transaksipenerimaanbarangheader_ibfk_1` FOREIGN KEY (`WhsIdf`) REFERENCES `masterwarehouse` (`WhsPK`),
  ADD CONSTRAINT `transaksipenerimaanbarangheader_ibfk_2` FOREIGN KEY (`TrxInSuppIdf`) REFERENCES `mastersupplier` (`SupplierPK`);

--
-- Ketidakleluasaan untuk tabel `transaksipengeluaranbarangdetail`
--
ALTER TABLE `transaksipengeluaranbarangdetail`
  ADD CONSTRAINT `transaksipengeluaranbarangdetail_ibfk_1` FOREIGN KEY (`TrxOutIDF`) REFERENCES `transaksipengeluaranbarangheader` (`TrxOutPK`),
  ADD CONSTRAINT `transaksipengeluaranbarangdetail_ibfk_2` FOREIGN KEY (`TrxOutDProductIdf`) REFERENCES `masterproduct` (`ProductPK`);

--
-- Ketidakleluasaan untuk tabel `transaksipengeluaranbarangheader`
--
ALTER TABLE `transaksipengeluaranbarangheader`
  ADD CONSTRAINT `transaksipengeluaranbarangheader_ibfk_1` FOREIGN KEY (`WhsIdf`) REFERENCES `masterwarehouse` (`WhsPK`),
  ADD CONSTRAINT `transaksipengeluaranbarangheader_ibfk_2` FOREIGN KEY (`TrxOutSuppIdf`) REFERENCES `mastersupplier` (`SupplierPK`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
