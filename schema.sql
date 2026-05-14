-- ====================== 1. 创建数据库 ======================
CREATE DATABASE IF NOT EXISTS jhb_platform DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE jhb_platform;

-- ====================== 2. 用户表 sys_user ======================
CREATE TABLE IF NOT EXISTS sys_user (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '用户ID',
    user_name VARCHAR(50) NOT NULL COMMENT '车主姓名',
    phone VARCHAR(20) NOT NULL COMMENT '手机号',
    plate_number VARCHAR(20) NOT NULL COMMENT '虚拟车牌',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    UNIQUE KEY uk_plate (plate_number),
    UNIQUE KEY uk_phone (phone)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 初始化40个模拟车主用户
INSERT INTO sys_user (user_name, phone, plate_number) VALUES
('车主01','13800000001','粤A00001'),('车主02','13800000002','粤A00002'),('车主03','13800000003','粤A00003'),('车主04','13800000004','粤A00004'),('车主05','13800000005','粤A00005'),
('车主06','13800000006','粤A00006'),('车主07','13800000007','粤A00007'),('车主08','13800000008','粤A00008'),('车主09','13800000009','粤A00009'),('车主10','13800000010','粤A00010'),
('车主11','13800000011','粤A00011'),('车主12','13800000012','粤A00012'),('车主13','13800000013','粤A00013'),('车主14','13800000014','粤A00014'),('车主15','13800000015','粤A00015'),
('车主16','13800000016','粤A00016'),('车主17','13800000017','粤A00017'),('车主18','13800000018','粤A00018'),('车主19','13800000019','粤A00019'),('车主20','13800000020','粤A00020'),
('车主21','13800000021','粤A00021'),('车主22','13800000022','粤A00022'),('车主23','13800000023','粤A00023'),('车主24','13800000024','粤A00024'),('车主25','13800000025','粤A00025'),
('车主26','13800000026','粤A00026'),('车主27','13800000027','粤A00027'),('车主28','13800000028','粤A00028'),('车主29','13800000029','粤A00029'),('车主30','13800000030','粤A00030'),
('车主31','13800000031','粤A00031'),('车主32','13800000032','粤A00032'),('车主33','13800000033','粤A00033'),('车主34','13800000034','粤A00034'),('车主35','13800000035','粤A00035'),
('车主36','13800000036','粤A00036'),('车主37','13800000037','粤A00037'),('车主38','13800000038','粤A00038'),('车主39','13800000039','粤A00039'),('车主40','13800000040','粤A00040');

-- ====================== 3. 通行里程表 traffic_record ======================
CREATE TABLE IF NOT EXISTS traffic_record (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '记录ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    plate_number VARCHAR(20) NOT NULL COMMENT '车牌',
    mileage DECIMAL(10,2) NOT NULL COMMENT '行驶里程(公里)',
    traffic_time DATETIME NOT NULL COMMENT '通行时间',
    is_calculate TINYINT DEFAULT 0 COMMENT '是否核算积分 0=未核算 1=已核算',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='通行里程记录表';

-- ====================== 4. 积分规则表 point_rule ======================
CREATE TABLE IF NOT EXISTS point_rule (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT '规则ID',
    mile_per_point INT NOT NULL COMMENT '每多少公里',
    point_value INT NOT NULL COMMENT '对应积分',
    is_default TINYINT DEFAULT 1 COMMENT '是否默认规则 1=是',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='积分规则表';

-- 初始化默认规则：100公里 = 10积分
INSERT INTO point_rule (mile_per_point, point_value) VALUES (100, 10);

-- ====================== 5. 用户积分余额表 user_point ======================
CREATE TABLE IF NOT EXISTS user_point (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    total_point INT DEFAULT 0 COMMENT '总积分余额',
    update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY uk_user_id (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户积分余额表';

-- ====================== 6. 积分变动日志表 point_log ======================
CREATE TABLE IF NOT EXISTS point_log (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '日志ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    change_point INT NOT NULL COMMENT '变动积分（正数增加 负数减少）',
    change_type VARCHAR(20) NOT NULL COMMENT '变动类型：里程入账/兑换优惠券/商品消费',
    relation_id BIGINT DEFAULT 0 COMMENT '关联ID（里程记录/优惠券/订单ID）',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='积分变动日志表';

-- ====================== 7. 优惠券配置表 coupon ======================
CREATE TABLE IF NOT EXISTS coupon (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT '优惠券ID',
    coupon_name VARCHAR(50) NOT NULL COMMENT '优惠券名称',
    discount_amount INT NOT NULL COMMENT '抵扣金额/积分',
    need_point INT NOT NULL COMMENT '兑换所需积分',
    stock INT DEFAULT 999 COMMENT '库存',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='优惠券配置表';

-- 初始化3种优惠券
INSERT INTO coupon (coupon_name, discount_amount, need_point) VALUES
('10元优惠券',10,50),('20元优惠券',20,100),('50元优惠券',50,200);

-- ====================== 8. 用户持有优惠券表 user_coupon ======================
CREATE TABLE IF NOT EXISTS user_coupon (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    coupon_id INT NOT NULL COMMENT '优惠券ID',
    status TINYINT DEFAULT 0 COMMENT '状态 0=未使用 1=已使用 2=已过期',
    exchange_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '兑换时间',
    use_time DATETIME NULL COMMENT '使用时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户优惠券表';

-- ====================== 9. 商品表 goods ======================
CREATE TABLE IF NOT EXISTS goods (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT '商品ID',
    goods_name VARCHAR(100) NOT NULL COMMENT '商品名称',
    need_point INT NOT NULL COMMENT '所需积分',
    stock INT DEFAULT 999 COMMENT '库存',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';

-- 初始化10个模拟商品
INSERT INTO goods (goods_name, need_point) VALUES
('车载充电器',50),('手机支架',80),('停车号码牌',30),('车载香薰',100),('玻璃水',40),
('洗车毛巾',20),('车载垃圾桶',60),('防滑垫',40),('数据线',70),('应急补胎液',150);

-- ====================== 10. 订单表 order_info ======================
CREATE TABLE IF NOT EXISTS order_info (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '订单ID',
    order_no VARCHAR(50) NOT NULL COMMENT '订单编号',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    goods_id INT NOT NULL COMMENT '商品ID',
    coupon_id INT DEFAULT 0 COMMENT '使用优惠券ID 0=未使用',
    use_point INT NOT NULL COMMENT '消耗积分',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '下单时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';
