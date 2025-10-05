
INSERT INTO subscriptions (id, user_id, plan, created_at, updated_at) VALUES
  (gen_random_uuid(), 34, 'basic', NOW() - INTERVAL '20 days', NOW() - INTERVAL '20 days'),
  (gen_random_uuid(), 34, 'pro', NOW() - INTERVAL '15 days', NOW() - INTERVAL '15 days'),
  (gen_random_uuid(), 34, 'enterprise', NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days'),
  (gen_random_uuid(), 34, 'pro', NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days'),
  (gen_random_uuid(), 34, 'basic', NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days');


INSERT INTO invoices (id, amount, status, created_at, updated_at) VALUES
  (gen_random_uuid(), 19.99, 'paid', NOW() - INTERVAL '19 days', NOW() - INTERVAL '19 days'),
  (gen_random_uuid(), 49.99, 'pending', NOW() - INTERVAL '18 days', NOW() - INTERVAL '18 days'),
  (gen_random_uuid(), 99.99, 'paid', NOW() - INTERVAL '17 days', NOW() - INTERVAL '17 days'),
  (gen_random_uuid(), 29.99, 'failed', NOW() - INTERVAL '16 days', NOW() - INTERVAL '16 days'),
  (gen_random_uuid(), 59.99, 'paid', NOW() - INTERVAL '15 days', NOW() - INTERVAL '15 days'),
  (gen_random_uuid(), 39.99, 'pending', NOW() - INTERVAL '14 days', NOW() - INTERVAL '14 days'),
  (gen_random_uuid(), 79.99, 'paid', NOW() - INTERVAL '13 days', NOW() - INTERVAL '13 days'),
  (gen_random_uuid(), 24.99, 'failed', NOW() - INTERVAL '12 days', NOW() - INTERVAL '12 days'),
  (gen_random_uuid(), 49.99, 'paid', NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days'),
  (gen_random_uuid(), 89.99, 'pending', NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days');
