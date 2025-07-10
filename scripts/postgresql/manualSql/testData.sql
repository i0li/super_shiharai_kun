-- ※ ユーザーを少なくとも1名は作成している必要がある

-- 10万件のテスト請求書データ作成
INSERT INTO public.invoices (
  user_id,
  issue_date,
  payment_amount,
  fee,
  fee_rate,
  tax_amount,
  tax_rate,
  total_amount,
  payment_due_date,
  created_at,
  updated_at
)
SELECT
  1,
  CURRENT_DATE - (i % 365), -- 最大1年前
  10000,
  500,
  0.04,
  800,
  0.10,
  12000, 
  CURRENT_DATE + (i % 365 + 1), -- 最大1年後
  NOW(),
  NOW()
FROM generate_series(1, 100000) AS s(i)
;
