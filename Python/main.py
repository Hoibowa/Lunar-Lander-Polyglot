import time

def get_valid_thrust(fuel):
    """燃料入力のバリデーション（ここがハックの技術点）"""
    while True:
        try:
            line = input(f"噴射する燃料を入力 (0-{min(10, fuel)}): ").strip()
            if not line: # 空エンターは0扱い
                return 0
            
            thrust = int(line)
            if 0 <= thrust <= 10:
                if thrust <= fuel:
                    return thrust
                else:
                    print(f"燃料が足りません（残り: {fuel}）")
            else:
                print("0から10の間で入力してください。")
        except ValueError:
            print("無効な入力です。数値を入力してください。")

def run_lunar_lander():
    while True:
        altitude, velocity, fuel = 100.0, 0.0, 50
        gravity, delta_t = 1.62, 1.0

        print("\n" + "="*40)
        print("--- LUNAR LANDER: MISSION START ---")
        print("="*40)
        
        while altitude > 0:
            print(f"高度: {altitude:6.2f}m | 速度: {velocity:6.2f}m/s | 燃料: {fuel:3}")
            
            thrust = get_valid_thrust(fuel)
            
            # 物理演算
            fuel -= thrust
            acceleration = gravity - (thrust * 0.4)
            velocity += acceleration * delta_t
            altitude -= velocity * delta_t

        print(f"\n--- タッチダウン！ 最終速度: {velocity:.2f} m/s ---")

        if velocity <= 5.0:
            print("【成功】完璧な着陸です。")
        elif velocity <= 10.0:
            print("【警告】機体中破。")
        else:
            print("【失敗】月面にクレーターを追加しました。")

        # リトライ判定
        while True:
            choice = input("\nもう一度挑戦しますか？ (y/n): ").strip().lower()
            if choice in ['y', 'yes']: break
            if choice in ['n', 'no', '']: return
            print("y か n で入力してください。")

if __name__ == "__main__":
    run_lunar_lander()