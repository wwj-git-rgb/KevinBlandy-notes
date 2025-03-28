-----------------------------
BeanCopier
-----------------------------
	# Spring 基于 CGLIB 实现的 Bean 属性 COPY

	# BeanCopier

		public static BeanCopier create(Class source, Class target, boolean useConverter)

			* 拷贝名称和类型都相同的属性。

			* source		源类型
			* target		目标类型
			* useConverter	是否使用自定义的转换器
				* true: 如果源对象和目标对象中同名属性的类型不一样，就会使用 copy 方法中的第三个参数 （Converter） 进行类型转换。
					* 如果设置为 true 则 copy 方法必须要传递第三个参数。
					* 如果返回的值不和目标方法中属性类型相匹配，则抛出异常。

				* false: 不适用转换器，它会对拷贝的对象和被拷贝的对象的类型进行判断，如果类型不同就不会拷贝。


		abstract public void copy(Object from, Object to, Converter converter)
	
	# BeanCopierKey
		public Object newInstance(String source, String target, boolean useConverter);
	

	# Converter
		Object convert(Object value, Class target, Object context);
			* value		是源对象属性的值
			* target	是目标对象同名属性的类型对象（Class）
			* context	源类属性的 setter 方法名称（String）

-----------------------------
BeanCopier - demo
-----------------------------
	# 相同类型属性 Copy
		class Foo {
			private Long id;
			private String title;
			private Instant createAt;
		}

		class Bar {
			private Long id;
			private String title;
			private Instant createAt;
		}

		public class MaiTest {

			public static void main(String[] args) throws Throwable {
			
				BeanCopier beanCopier = BeanCopier.create(Foo.class, Bar.class, false);
				
				var bar = new Bar();
				
				beanCopier.copy(Foo.builder()
						.id(1000L)
						.title("我是FOO")
						.createAt(Instant.now())
						.build(), bar, null);
				
				// Bar(id=1000, title=我是FOO, createAt=2025-01-22T06:57:01.042756600Z)
				System.out.println(bar);
			}
		}
	
	# 包含不相同类型的属性 Copy

		class Foo {
			private Long id;
			private String title;
			private Instant createAt;
		}


		class Bar {
			private Long id;
			private String title;
			private Long createAt;  // 这里是 Long
		}

		public class MaiTest {

			
			public static void main(String[] args) throws Throwable {
			
				BeanCopier beanCopier = BeanCopier.create(Foo.class, Bar.class, true);
				
				var bar = new Bar();
				
				beanCopier.copy(Foo.builder()
						.id(1000L)
						.title("我是FOO")
						.createAt(Instant.now())
						.build(), bar, (value, target, content) -> {
							// 源属性值是 Instant 且目标属性值类型是 Long
							if (value instanceof Instant time && ((Class<?>)target).isAssignableFrom(Long.class)) {
								return time.toEpochMilli();
							}
							return value;
						});
				
				// Bar(id=1000, title=我是FOO, createAt=1737530473591)
				System.out.println(bar);
			}
		}


-----------------------------
ModelMapper
-----------------------------

import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentMap;

import org.springframework.cglib.beans.BeanCopier;

import lombok.extern.slf4j.Slf4j;

@Slf4j
public class ModelMapper {

	private ModelMapper() {
		throw new RuntimeException();
	}

	static final ConcurrentMap<String, BeanCopier> CONCURRENT_MAP = new ConcurrentHashMap<>();

	/**
	 * 复制 src 对象的数学到 dest，只 copy 同名且类型一致的属性
	 * 
	 * @param <S>
	 * @param <T>
	 * @param src
	 * @param dest
	 * @return
	 */
	public static <S, T> T copy(S src, T dest) {

		Objects.requireNonNull(src);
		Objects.requireNonNull(dest);

		Class<?> srcClazz = src.getClass();
		Class<?> destClazz = dest.getClass();

		CONCURRENT_MAP.compute(key(srcClazz, destClazz), (k, v) -> {
			if (v == null) {
				log.info("Create BeanCopier: {} -> {}", srcClazz, destClazz);
				v = BeanCopier.create(srcClazz, destClazz, false);
			}
			return v;
		}).copy(src, dest, null);
		;

		return dest;
	}

	private static String key(Class<?> src, Class<?> dest) {
		return src.getName() + "_" + dest.getName();
	}
}








